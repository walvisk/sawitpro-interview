package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/service/auth"
	"github.com/SawitProRecruitment/UserService/service/user"
	"github.com/SawitProRecruitment/UserService/service/userlog"
	"github.com/SawitProRecruitment/UserService/utils"
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

	responseBody := strings.TrimSpace(rec.Body.String())
	expectedResponseBody := `{"id":1}`
	if responseBody != expectedResponseBody {
		t.Fatalf("Expected response body: %s, got: %s", expectedResponseBody, responseBody)
	}
}

func TestLogin(t *testing.T) {
	dummyPwd := "d0lor@Ipsum"
	hashedPwd, err := utils.HashPassword(dummyPwd)
	if err != nil {
		t.FailNow()
	}
	dummyUsr := &repository.User{
		ID:       int64(1),
		Password: hashedPwd,
		Phone:    "8133402912",
	}

	jsonPayload := `{"phone":"+628133402912", "password": "d0lor@Ipsum"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userServiceMock := user.NewMockService(ctrl)
	authServiceMock := auth.NewMockService(ctrl)
	userLogServiceMock := userlog.NewMockService(ctrl)
	opts := NewServerOptions{
		UserService:    userServiceMock,
		AuthService:    authServiceMock,
		UserLogService: userLogServiceMock,
	}
	userServiceMock.EXPECT().FindUserByPhone(c.Request().Context(), "+628133402912").Return(dummyUsr, nil)
	authServiceMock.EXPECT().AuthenticateUserPassword(dummyUsr, dummyPwd).Return(nil)
	authServiceMock.EXPECT().GenerateJWT().Return("dummyToken", nil)
	userLogServiceMock.EXPECT().CreateUserLog(c.Request().Context(), dummyUsr).Return(nil)
	h := NewServer(opts)
	h.Login(c)

	if !gomock.Eq(http.StatusOK).Matches(rec.Code) {
		t.Fatalf("Expected HTTP status %d, got %d", http.StatusOK, rec.Code)
	}

	responseBody := strings.TrimSpace(rec.Body.String())
	expectedResponseBody := `{"id":1,"token":"dummyToken"}`
	if responseBody != expectedResponseBody {
		t.Fatalf("Expected response body: %s, got: %s", expectedResponseBody, responseBody)
	}
}

func TestProfile(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/profile/123", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/profile/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")
	c.Request().Header.Set("Authorization", "Bearer valid_token")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userServiceMock := user.NewMockService(ctrl)
	authServiceMock := auth.NewMockService(ctrl)

	opts := NewServerOptions{
		UserService: userServiceMock,
		AuthService: authServiceMock,
	}
	h := NewServer(opts)

	authServiceMock.EXPECT().ValidateJWT(gomock.Any()).Return(nil)
	userServiceMock.EXPECT().FindUserByID(gomock.Any(), int64(123)).Return(&repository.User{
		FullName: "Dolor Ipsum Sit Amet",
		Phone:    "+62123456789",
	}, nil)

	err := h.Profile(c, int64(123))
	if err != nil {
		t.Fatalf("Error handling request: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected HTTP status %d, got %d", http.StatusOK, rec.Code)
	}

	responseBody := strings.TrimSpace(rec.Body.String())
	expectedResponseBody := `{"full_name":"Dolor Ipsum Sit Amet","phone":"+62123456789"}`
	if responseBody != expectedResponseBody {
		t.Fatalf("Expected response body: %s, got: %s", expectedResponseBody, rec.Body.String())
	}
}

func TestUpdateUser(t *testing.T) {
	jsonPayload := `{"full_name": "Dolor Ipsum New", "phone": "+621234567810"}`

	e := echo.New()

	req := httptest.NewRequest(http.MethodPatch, "/profile/123", strings.NewReader(jsonPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/profile/:id")
	c.SetParamNames("id")
	c.SetParamValues("123")
	c.Request().Header.Set("Authorization", "Bearer valid_token")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userServiceMock := user.NewMockService(ctrl)
	authServiceMock := auth.NewMockService(ctrl)

	opts := NewServerOptions{
		UserService: userServiceMock,
		AuthService: authServiceMock,
	}
	h := NewServer(opts)

	authServiceMock.EXPECT().ValidateJWT(gomock.Any()).Return(nil)
	userServiceMock.EXPECT().FindUserByID(gomock.Any(), int64(123)).Return(&repository.User{
		FullName: "Dolor Ipsum Old",
		Phone:    "+62123456789",
	}, nil)
	userServiceMock.EXPECT().FindUserByPhone(gomock.Any(), "+621234567810").Return(nil, nil)
	userServiceMock.EXPECT().UpdateUser(gomock.Any(), gomock.Any(), "Dolor Ipsum New", "+621234567810").Return(nil)

	err := h.UpdateUser(c, int64(123))
	if err != nil {
		t.Fatalf("Error handling request: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected HTTP status %d, got %d", http.StatusOK, rec.Code)
	}
}
