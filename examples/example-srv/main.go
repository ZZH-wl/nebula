package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"example-srv/handler"
	"example-srv/subscriber"

	hello "example-srv/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("nebula.core.srv.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), new(subscriber.Hello))

	// Register Function as Subscriber
	micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
