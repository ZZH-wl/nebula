package client

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	hello "path/to/service/proto/hello"
)

type helloKey struct{}

// FromContext retrieves the client from the Context
func HelloFromContext(ctx context.Context) (hello.HelloService, bool) {
	c, ok := ctx.Value(helloKey{}).(hello.HelloService)
	return c, ok
}

// Client returns a wrapper for the HelloClient
func HelloWrapper(service micro.Service) server.HandlerWrapper {
	client := hello.NewHelloService("go.micro.srv.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, helloKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
