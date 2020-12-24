// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/comments.proto

package comments

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Comments service

func NewCommentsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Comments service

type CommentsService interface {
	Store(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type commentsService struct {
	c    client.Client
	name string
}

func NewCommentsService(name string, c client.Client) CommentsService {
	return &commentsService{
		c:    c,
		name: name,
	}
}

func (c *commentsService) Store(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "Comments.Store", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentsService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Comments.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Comments service

type CommentsHandler interface {
	Store(context.Context, *StoreRequest, *StoreResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterCommentsHandler(s server.Server, hdlr CommentsHandler, opts ...server.HandlerOption) error {
	type comments interface {
		Store(ctx context.Context, in *StoreRequest, out *StoreResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type Comments struct {
		comments
	}
	h := &commentsHandler{hdlr}
	return s.Handle(s.NewHandler(&Comments{h}, opts...))
}

type commentsHandler struct {
	CommentsHandler
}

func (h *commentsHandler) Store(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.CommentsHandler.Store(ctx, in, out)
}

func (h *commentsHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.CommentsHandler.Get(ctx, in, out)
}