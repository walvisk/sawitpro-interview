package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) CreateUser(c echo.Context) error {
	var payload generated.CreateUserRequest
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: "invalid json format",
		})
	}

	uv := UserValidator{
		FullName: payload.FullName,
		Phone:    payload.Phone,
		Password: payload.Password,
	}
	uv.Validate()
	if uv.HasError() {
		return c.JSON(http.StatusBadRequest, uv.UserError)
	}

	id, err := s.UserService.RegisterUser(c.Request().Context(), payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generated.ErrorResponse{
			Kind:    "BadRequest",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, generated.CreateUserResponse{Id: id})
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

	return c.JSON(http.StatusOK, generated.LoginResponse{
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
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, generated.ProfileResponse{
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
