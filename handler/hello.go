package handler

import (
	"api-sketch/biz"
	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
	Hello gin.HandlerFunc `method:"get" path:"/hello"`
}

func NewHelloHandler(hb biz.HelloBiz) *HelloHandler {
	h := &helloHandler{
		hb: hb,
	}
	return &HelloHandler{
		Hello: h.Hello,
	}
}

type helloHandler struct {
	hb biz.HelloBiz
}

func (h *helloHandler) Hello(c *gin.Context) {
	res := h.hb.Hello()
	c.JSON(0, res)
}
