package handler

import (
	"net/http"

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
