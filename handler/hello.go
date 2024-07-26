package handler

import (
	"api-sketch/biz"
	"api-sketch/proto"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	Hello        gin.HandlerFunc `method:"post" path:"/hello"`
	Ping         gin.HandlerFunc `method:"get" path:"/ping"`
	HelloService gin.HandlerFunc `method:"get" path:"/hello-service"`
}

func NewHelloHandler(hb biz.HelloBiz) *HelloHandler {
	h := &helloHandler{
		hb: hb,
		hs: proto.NewHelloService(proto.ServiceName, proto.Service.Client()),
	}
	return &HelloHandler{
		Hello:        h.Hello,
		Ping:         h.Ping,
		HelloService: h.HelloService,
	}
}

type helloHandler struct {
	hb biz.HelloBiz
	hs proto.HelloService
}

func (h *helloHandler) Hello(c *gin.Context) {
	var args struct {
		Name string `json:"name"`
	}
	if err := c.Bind(&args); handleError(c, err) {
		return
	}

	ctx, cancel := createContext()
	defer cancel()

	if err := h.hb.Hello(ctx, args.Name); handleError(c, err) {
		return
	}
	response(c, nil)
}

func (h *helloHandler) Ping(c *gin.Context) {
	ctx, cancel := createContext()
	defer cancel()

	user, _ := c.Get("user")
	name := c.Query("name")
	id, err := h.hb.Ping(ctx, name)
	if handleError(c, err) {
		return
	}
	response(c, map[string]any{
		"user": user,
		"name": name,
		"id":   id,
	})
}

func (h *helloHandler) HelloService(c *gin.Context) {
	ctx, cancel := createContext()
	defer cancel()

	var args struct {
		Name string `form:"name"`
	}
	if err := c.BindQuery(&args); handleError(c, err) {
		return
	}

	res, err := h.hs.SayHello(ctx, &proto.HelloRequest{Name: args.Name})
	if handleError(c, err) {
		return
	}
	response(c, map[string]any{
		"res": res.Message,
	})
}
