// Code scaffolded by goctl. Safe to edit.
// goctl <no value>

package svc

import (
	"hello/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
