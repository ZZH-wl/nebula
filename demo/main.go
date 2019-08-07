package main

import (
	"github.com/Wall-js/nebula"
	hello "github.com/Wall-js/nebula/demo/grpc/service/hello"
	"github.com/Wall-js/nebula/demo/handler"
	"github.com/Wall-js/nebula/demo/subscriber"
	"github.com/micro/go-micro"
)

func main() {
	nebula.Service.Init(
		micro.Name("nebula.core.srv.hello"),
	)
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), new(subscriber.Hello))
	micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), subscriber.Handler)
	nebula.Run()
}
