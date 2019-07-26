package handler

import (
	"context"

	hello "example-fnc/proto/hello"
)

type Hello struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Hello) Call(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	rsp.Msg = "Hello " + req.Name
	return nil
}
