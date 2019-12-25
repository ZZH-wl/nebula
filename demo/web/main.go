package main

import (
	"github.com/Wall-js/nebula"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
)

type Say struct{}

func (s *Say) Anything(c *gin.Context) {
	log.Log("Received Say.Anything API request")
	c.JSON(200, map[string]string{
		"message": "Hi, this is the Greeter API",
	})
}

func main() {
	nebula.SetName("com.tradeany.api.greeter")
	router := gin.New()
	say := new(Say)
	router.GET("/greeter", say.Anything)
	router.GET("/greeter/test", say.Anything)
	// Register Handler
	nebula.Web.Handle("/", router)
	nebula.RunWeb()
}
