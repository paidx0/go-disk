package svc

import (
	"code/internal/config"
	"code/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Middleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		Middleware: middleware.NewMiddlewareMiddleware().Handle,
	}
}
