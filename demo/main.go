package main

import (
	"github.com/Wall-js/nebula"
	"github.com/micro/go-micro"
)

func main() {
	nebula.Run(func(service micro.Service) {
		service.Init(
			micro.Name("nebula.core.srv.hello"),
		)
	})
}
