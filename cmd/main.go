package main

import (
	"log"
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/service/auth"
	"github.com/SawitProRecruitment/UserService/service/user"
	userLog "github.com/SawitProRecruitment/UserService/service/user_log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	userService := user.NewUserService(repo)
	authServiceOpts := auth.AuthServiceOpts{
		Repository:     repo,
		PrivateKeyFile: os.Getenv("PRIVATE_KEY_FILE"),
		PublicKeyFile:  os.Getenv("PUBLIC_KEY_FILE"),
	}

	authService, err := auth.NewAuthService(authServiceOpts)
	if err != nil {
		log.Fatal(err)
	}

	userLogService := userLog.NewUserLogService(repo)

	opts := handler.NewServerOptions{
		UserService:    userService,
		AuthService:    authService,
		UserLogService: userLogService,
	}
	return handler.NewServer(opts)
}
