package svc

import "github.com/fyerfyer/go-zero-ecommerce/ecommerce/comment/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
