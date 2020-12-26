package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/tags/model"
	tags "github.com/emadghaffari/micro-blog/tags/proto"
)

// Tags struct
type Tags struct {
	m model.Model
}

// NewTags return Tags pointer
func NewTags() *Tags {
	return &Tags{m: model.New("tags")}
}

// Add is a single request handler called via client.Call or the generated client code
func (e *Tags) Add(ctx context.Context, req *tags.AddRequest, rsp *tags.AddResponse) error {
	log.Warn("ADDED NEW TAG FROM POST: %d", req.Title)
	return nil
}

// List is a single request handler called via client.Call or the generated client code
func (e *Tags) List(ctx context.Context, req *tags.ListRequest, rsp *tags.ListResponse) error {

	return nil
}

// Remove is a single request handler called via client.Call or the generated client code
func (e *Tags) Remove(ctx context.Context, req *tags.RemoveRequest, rsp *tags.RemoveResponse) error {

	return nil
}

// Update is a single request handler called via client.Call or the generated client code
func (e *Tags) Update(ctx context.Context, req *tags.UpdateRequest, rsp *tags.UpdateResponse) error {

	return nil
}
