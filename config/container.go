package config

import (
    "github.com/douglasfsti/golang-shortener-api/pkg/shortner"
)

type Container interface {
	// shortner
	GetShortnerService() shortner.Service
}

func NewContainer(service shortner.Service) Container {
	return &container{
		ShortnerService: service,
	}
}

type container struct {
	ShortnerService shortner.Service
}

// shortner interfaces
func (c *container) GetShortnerService() shortner.Service {
	return c.ShortnerService
}
