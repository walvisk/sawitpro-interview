package handler

import (
	"github.com/SawitProRecruitment/UserService/service/auth"
	"github.com/SawitProRecruitment/UserService/service/user"
)

type Server struct {
	UserService user.Service
	AuthService auth.Service
}

type NewServerOptions struct {
	UserService user.Service
	AuthService auth.Service
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserService: opts.UserService,
		AuthService: opts.AuthService,
	}
}
