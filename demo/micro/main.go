package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

type Greeter struct{}

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.31.131:8500",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("helloworld"))

	service.Init()

	if err := service.Run(); err != nil {
		fmt.Println("failed to run a service: ", err)
	}
}
