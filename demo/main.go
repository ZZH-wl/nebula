package main

import (
	"github.com/Wall-js/nebula"
	hello "github.com/Wall-js/nebula/demo/grpc/service/hello"
	"github.com/Wall-js/nebula/demo/handler"
	"github.com/Wall-js/nebula/demo/subscriber"
	"github.com/micro/go-micro"
)

func main() {
	nebula.Run(func(service micro.Service) {
		service.Init(
			micro.Name("nebula.core.srv.hello"),
			micro.Version("latest"),
		)
		// Register Handler
		hello.RegisterHelloHandler(service.Server(), new(handler.Hello))

		// Register Struct as Subscriber
		micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), new(subscriber.Hello))

		// Register Function as Subscriber
		micro.RegisterSubscriber("nebula.core.srv.hello", service.Server(), subscriber.Handler)

	})
}
