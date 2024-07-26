// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package proto

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for HelloService service

func NewHelloServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for HelloService service

type HelloService interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) SayHello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	req := c.c.NewRequest(c.name, "HelloService.SayHello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloService service

type HelloServiceHandler interface {
	SayHello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterHelloServiceHandler(s server.Server, hdlr HelloServiceHandler, opts ...server.HandlerOption) error {
	type helloService interface {
		SayHello(ctx context.Context, in *HelloRequest, out *HelloResponse) error
	}
	type HelloService struct {
		helloService
	}
	h := &helloServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HelloService{h}, opts...))
}

type helloServiceHandler struct {
	HelloServiceHandler
}

func (h *helloServiceHandler) SayHello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.HelloServiceHandler.SayHello(ctx, in, out)
}
