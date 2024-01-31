package handler

import (
	"github.com/SawitProRecruitment/UserService/service/auth"
	"github.com/SawitProRecruitment/UserService/service/user"
	"github.com/SawitProRecruitment/UserService/service/userlog"
)

type Server struct {
	UserService    user.Service
	AuthService    auth.Service
	UserLogService userlog.Service
}

type NewServerOptions struct {
	UserService    user.Service
	AuthService    auth.Service
	UserLogService userlog.Service
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserService:    opts.UserService,
		AuthService:    opts.AuthService,
		UserLogService: opts.UserLogService,
	}
}
