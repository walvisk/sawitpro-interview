package handler

import (
	"github.com/SawitProRecruitment/UserService/service/auth"
	"github.com/SawitProRecruitment/UserService/service/user"
	userLog "github.com/SawitProRecruitment/UserService/service/user_log"
)

type Server struct {
	UserService    user.Service
	AuthService    auth.Service
	UserLogService userLog.Service
}

type NewServerOptions struct {
	UserService    user.Service
	AuthService    auth.Service
	UserLogService userLog.Service
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserService:    opts.UserService,
		AuthService:    opts.AuthService,
		UserLogService: opts.UserLogService,
	}
}
