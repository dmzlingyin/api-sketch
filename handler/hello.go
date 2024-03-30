package handler

import "github.com/gin-gonic/gin"

type HelloHandler struct {
	Hello gin.HandlerFunc `method:"get" path:"/hello"`
}

func NewHelloHandler() *HelloHandler {
	h := &helloHandler{}
	return &HelloHandler{
		Hello: h.Hello,
	}
}

type helloHandler struct {
}

func (h *helloHandler) Hello(c *gin.Context) {
	c.JSON(0, "world")
}
