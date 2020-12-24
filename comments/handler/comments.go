package handler

import (
	"context"
	"fmt"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/comments/model"
	comments "github.com/emadghaffari/micro-blog/comments/proto"
)

type Comments struct {
	store model.Model
}

// NewComments func
func NewComments() *Comments {
	return &Comments{
		store: model.New("comments"),
	}
}

// Store is a single request handler called via client.Call or the generated client code
func (e *Comments) Store(ctx context.Context, req *comments.StoreRequest, rsp *comments.StoreResponse) error {
	log.Info("Received Comments.Call request")
	err := e.store.Save(req)
	if err != nil {
		return err
	}
	rsp.Msg = "Stored New Comment by [" + fmt.Sprintf("%d", req.UserId) + "] for: [" + fmt.Sprintf("%d", req.PostId) + "] post"
	return nil
}

// Get is a single request handler called via client.Call or the generated client code
func (e *Comments) Get(ctx context.Context, req *comments.GetRequest, rsp *comments.GetResponse) error {
	log.Info("Received Comments.Call request")
	return e.store.List(req, &rsp.Comments)
}
