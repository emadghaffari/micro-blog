// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/posts.proto

package posts

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

// Api Endpoints for Posts service

func NewPostsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Posts service

type PostsService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Store(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type postsService struct {
	c    client.Client
	name string
}

func NewPostsService(name string, c client.Client) PostsService {
	return &postsService{
		c:    c,
		name: name,
	}
}

func (c *postsService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Posts.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Store(ctx context.Context, in *StoreRequest, opts ...client.CallOption) (*StoreResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Store", in)
	out := new(StoreResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postsService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "Posts.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Posts service

type PostsHandler interface {
	Call(context.Context, *Request, *Response) error
	Store(context.Context, *StoreRequest, *StoreResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterPostsHandler(s server.Server, hdlr PostsHandler, opts ...server.HandlerOption) error {
	type posts interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Store(ctx context.Context, in *StoreRequest, out *StoreResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type Posts struct {
		posts
	}
	h := &postsHandler{hdlr}
	return s.Handle(s.NewHandler(&Posts{h}, opts...))
}

type postsHandler struct {
	PostsHandler
}

func (h *postsHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.PostsHandler.Call(ctx, in, out)
}

func (h *postsHandler) Store(ctx context.Context, in *StoreRequest, out *StoreResponse) error {
	return h.PostsHandler.Store(ctx, in, out)
}

func (h *postsHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.PostsHandler.Get(ctx, in, out)
}
