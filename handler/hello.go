package handler

import (
	"api-sketch/biz"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	Hello gin.HandlerFunc `method:"post" path:"/hello"`
	Ping  gin.HandlerFunc `method:"get" path:"/ping"`
}

func NewHelloHandler(hb biz.HelloBiz) *HelloHandler {
	h := &helloHandler{
		hb: hb,
	}
	return &HelloHandler{
		Hello: h.Hello,
		Ping:  h.Ping,
	}
}

type helloHandler struct {
	hb biz.HelloBiz
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
