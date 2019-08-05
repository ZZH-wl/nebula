package main

import (
	"github.com/Wall-js/nebula"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"log"
)

type Say struct{}

func (s *Say) Anything(c *gin.Context) {
	log.Print("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main() {
	nebula.Web.Init(
		web.Name("go.micro.api.greeter"),
	)

	router := gin.New()
	say := new(Say)
	router.GET("/greeter", say.Anything)

	// Register Handler
	nebula.Web.Handle("/", router)
	nebula.RunWeb()
}

//func main() {
//
//	service := web.NewService(
//		web.Name("go.micro.api.greeter"),
//	)
//
//	reg := etcdv3.NewRegistry(func(op *registry.Options) {
//		op.Addrs = []string{"http://localhost:2379"}
//		//op.Addrs = Conf.Get("registryAddr").StringSlice([]string{"localhost:2379"})
//	})
//
//	if err := service.Init(
//		web.Registry(reg),
//	); err != nil {
//		log.Fatal(err)
//	}
//	// setup Greeter Server Client
//	service.Options()
//	// Create RESTful handler (using Gin)
//	say := new(Say)
//	router := gin.Default()
//	router.GET("/greeter", say.Anything)
//
//	// Register Handler
//	service.Handle("/", router)
//
//	// Run server
//	if err := service.Run(); err != nil {
//		log.Fatal(err)
//	}
//}