package handler

import (
	"context"

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
	return nil
}

// Get is a single request handler called via client.Call or the generated client code
func (e *Comments) Get(ctx context.Context, req *comments.StoreRequest, rsp *comments.StoreResponse) error {
	log.Info("Received Comments.Call request")
	return nil
}
