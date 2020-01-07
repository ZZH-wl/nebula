package main

import (
	hello "github.com/Wall-js/nebula/demo/grpc/service/hello"
	"github.com/Wall-js/nebula/demo/handler"
	"github.com/Wall-js/nebula/v2"
)

func main() {
	nebula.SetName("com.nebula.test")
	nebula.AddPrefix("test/")
	//nebula.SetConfigKey("nebula/nebula-core/latest")
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	nebula.Run()
}
