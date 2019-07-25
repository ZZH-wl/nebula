package nebula

import (
	"github.com/micro/go-micro"
	"testing"
)

func TestRun(t *testing.T) {
	Run(func(service *micro.Service) {

	})
}

func TestShunDown(t *testing.T) {
	ShutDown()
}
