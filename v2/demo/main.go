package main

import (
	"fmt"
	"github.com/Wall-js/nebula"
	hello "github.com/Wall-js/nebula/demo/grpc/service/hello"
	"github.com/Wall-js/nebula/demo/handler"
)

func main() {
	nebula.SetName("com.nebula.test")
	nebula.AddPrefix("/nebula/test")
	//nebula.SetConfigKey("nebula/nebula-core/latest")
	hello.RegisterHelloHandler(nebula.Service.Server(), new(handler.Hello))
	nebula.BeforeStart(
		func() { fmt.Println("666") },
	)
	nebula.Run()
}
