package handler

import (
	"github.com/SawitProRecruitment/UserService/service/user"
)

type Server struct {
	UserService user.Service
}

type NewServerOptions struct {
	UserService user.Service
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		UserService: opts.UserService,
	}
}
