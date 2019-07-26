package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"example-fnc/handler"
	"example-fnc/subscriber"
)

func main() {
	// New Service
	function := micro.NewFunction(
		micro.Name("nebula.core.fnc.hello"),
		micro.Version("latest"),
	)

	// Initialise function
	function.Init()

	// Register Handler
	function.Handle(new(handler.Hello))

	// Register Struct as Subscriber
	function.Subscribe("nebula.core.fnc.hello", new(subscriber.Hello))

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
