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
	ctx, cancel := createContext()
	defer cancel()

	if err := h.hb.Hello(ctx, "alice"); err != nil {
		c.JSON(1, nil)
	}
	c.JSON(0, map[string]any{
		"code": 0,
	})
}

func (h *helloHandler) Ping(c *gin.Context) {
	c.JSON(0, map[string]any{
		"message": "pong",
	})
}
