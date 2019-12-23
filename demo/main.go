package main

import (
	"github.com/Wall-js/nebula"
	hello "github.com/Wall-js/nebula/demo/grpc/service/hello"
	"github.com/Wall-js/nebula/demo/handler"
)

func main() {
	//nebula.Service.Init(
	//	micro.Name("nebula.core.srv.88"),
	//)
	nebula.SetName("com.nebula.test")
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	//micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), new(subscriber.Hello))
	//micro.RegisterSubscriber("nebula.core.srv.hello", nebula.Service.Server(), subscriber.Handler)
	nebula.Run()
}
