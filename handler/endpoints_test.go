package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/service/user"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func TestCreateUser(t *testing.T) {
	jsonPayload := `{"full_name": "Hei Hei", "phone":"+628133402912", "password": "d0lor@Ipsum"}`

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

	userServiceMock.EXPECT().FindUserByPhone(c.Request().Context(), "+628133402912").Return(nil, nil)
	userServiceMock.EXPECT().RegisterUser(c.Request().Context(), generated.CreateUserJSONRequestBody{
		FullName: "Hei Hei",
		Phone:    "+628133402912",
		Password: "d0lor@Ipsum",
	}).Return(int64(1), nil)
	h.CreateUser(c)

	if !gomock.Eq(http.StatusCreated).Matches(rec.Code) {
		t.Fatalf("Expected HTTP status %d, got %d", http.StatusCreated, rec.Code)
	}

}
