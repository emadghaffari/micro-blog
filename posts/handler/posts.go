package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	"github.com/emadghaffari/micro-blog/posts/model"
	posts "github.com/emadghaffari/micro-blog/posts/proto"
)

// Posts struct
type Posts struct {
	m model.Model
}

// NewPosts return Posts pointer
func NewPosts() *Posts {
	return &Posts{m: model.New("posts")}
}

// Call is a single request handler called via client.Call or the generated client code
func (p *Posts) Call(ctx context.Context, req *posts.Request, rsp *posts.Response) error {
	log.Info("Received Posts.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Store is a single request handler Store via client.Store or the generated client code
func (p *Posts) Store(ctx context.Context, req *posts.StoreRequest, rsp *posts.StoreResponse) error {
	log.Info("Received Posts.Store request")
	if err := p.m.Save(req); err != nil {
		log.Warn(err.Error())
		rsp.Msg = err.Error()
	}
	rsp.Msg = "stored " + req.Title + " by user: " + req.UserId
	return nil
}

// Get is a single request handler Get via client.Get or the generated client code
func (p *Posts) Get(ctx context.Context, req *posts.GetRequest, rsp *posts.GetResponse) error {
	log.Info("Received Posts.Store request")
	return p.m.List(req, &rsp.Posts)
}
