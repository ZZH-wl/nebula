package main

import (
	"github.com/Wall-js/nebula/demo/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://localhost:2379"}
	})
	// New Service
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.micro"),
		micro.Version("latest"),
	)
	// Initialise service
	service.Init()
	// Register Handler
	micro.RegisterHandler(service.Server(), new(handler.Hello))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
