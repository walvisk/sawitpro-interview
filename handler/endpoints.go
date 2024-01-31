package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateUser(c echo.Context) error {
	var payload generated.BaseUser
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	uv := UserValidator{
		FullName: payload.FullName,
		Phone:    payload.Phone,
		Password: payload.Password,
	}
	uv.ValidateFullName().ValidatePhone().ValidatePassword()
	if uv.HasError() {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "invalid request payload",
			Fields:  uv.Errors,
		})
	}

	checkUser, err := s.UserService.FindUserByPhone(c.Request().Context(), payload.Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Kind:    "InternalServerError",
			Message: err.Error(),
		})
	}
	if checkUser != nil {
		return c.JSON(http.StatusConflict, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "phone already existed",
		})
	}

	id, err := s.UserService.RegisterUser(c.Request().Context(), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, generated.BaseUser{Id: id})
}

func (s *Server) Login(c echo.Context) error {
	var payload generated.LoginJSONRequestBody
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "invalid json format",
		})
	}

	user, err := s.UserService.FindUserByPhone(c.Request().Context(), payload.Phone)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusConflict, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "user not found",
		})
	}

	err = s.AuthService.AuthenticateUserPassword(user, payload.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	token, err := s.AuthService.GenerateJWT()
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	err = s.UserLogService.CreateUserLog(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Kind:    "InternalServerError",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, generated.BaseUser{
		Id:    user.ID,
		Token: token,
	})
}

func (s *Server) Profile(c echo.Context, id int64) error {
	authToken, err := getToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, generated.ErrorResponse{
			Kind:    "Forbidden",
			Message: err.Error(),
		})
	}

	err = s.AuthService.ValidateJWT(authToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, generated.ErrorResponse{
			Kind:    "Forbidden",
			Message: err.Error(),
		})
	}

	user, err := s.UserService.FindUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Kind:    "InternalServerError",
			Message: err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, generated.BaseUser{
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

func (s *Server) UpdateUser(c echo.Context, id int64) error {
	authToken, err := getToken(c)
	if err != nil {
		return c.JSON(http.StatusForbidden, generated.ErrorResponse{
			Kind:    "Forbidden",
			Message: err.Error(),
		})
	}

	err = s.AuthService.ValidateJWT(authToken)
	if err != nil {
		return c.JSON(http.StatusForbidden, generated.ErrorResponse{
			Kind:    "Forbidden",
			Message: err.Error(),
		})
	}

	var payload generated.BaseUser
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "invalid json format",
		})
	}

	uv := UserValidator{
		FullName: payload.FullName,
		Phone:    payload.Phone,
	}
	if payload.FullName != "" {
		uv.ValidateFullName()
	}
	if payload.Phone != "" {
		uv.ValidatePhone()
	}
	if uv.HasError() {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "invalid request payload",
			Fields:  uv.Errors,
		})
	}

	user, err := s.UserService.FindUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Kind:    "InternalServerError",
			Message: err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "user not found",
		})
	}

	checkUser, err := s.UserService.FindUserByPhone(c.Request().Context(), payload.Phone)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, generated.ErrorResponse{
			Kind:    "InternalServerError",
			Message: err.Error(),
		})
	}
	if checkUser != nil {
		return c.JSON(http.StatusConflict, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "phone already existed",
		})
	}

	err = s.UserService.UpdateUser(c.Request().Context(), user, payload.Phone, payload.FullName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, generated.BaseUser{
		FullName: user.FullName,
		Phone:    user.Phone,
	})
}

func getToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return "", errors.New("invalid token")
	}

	return strings.Replace(authHeader, "Bearer ", "", -1), nil
}
