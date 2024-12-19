package svc

import "github.com/fyerfyer/go-zero-ecommerce/ecommerce/reply/rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
