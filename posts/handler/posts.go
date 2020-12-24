package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	posts "github.com/emadghaffari/micro-blog/posts/proto"
)

// Posts struct
type Posts struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Posts) Call(ctx context.Context, req *posts.Request, rsp *posts.Response) error {
	log.Info("Received Posts.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}
