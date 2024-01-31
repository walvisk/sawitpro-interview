package handler

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/service/user"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestCreateUser(t *testing.T) {
	jsonPayload := `{"full_name": "Hei Hei", "phone":"+6281334029", "password": "d0lor@Ipsum"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userServiceMock := user.NewMockService(ctrl)

	opts := NewServerOptions{
		UserService: userServiceMock,
	}
	h := NewServer(opts)
	h.CreateUser(c)
	log.Println(rec.Body)
	log.Println(rec.Code)
	gomock.Eq(http.StatusCreated).Matches(rec.Code)
}
