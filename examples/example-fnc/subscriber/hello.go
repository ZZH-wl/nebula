package subscriber

import (
	"context"

	"github.com/micro/go-micro/util/log"

	hello "example-fnc/proto/hello"
)

type Hello struct{}

func (e *Hello) Handle(ctx context.Context, msg *hello.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
