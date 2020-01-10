package nebula

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/web"
)

type Options struct {
	Version     string
	Conf        config.Config
	Service     micro.Service
	Web         web.Service
	appId       string
	serviceName string
	process     func()
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	opt := Options{
		Version:     "v1.17.4",
		Conf:        config.NewConfig(),
		Service:     micro.NewService(),
		Web:         web.NewService(),
		appId:       "",
		serviceName: "",
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}
