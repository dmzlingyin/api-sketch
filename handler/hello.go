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
	if err := c.Bind(&args); err != nil {
		c.JSON(200, nil)
		return
	}

	ctx, cancel := createContext()
	defer cancel()

	if err := h.hb.Hello(ctx, args.Name); err != nil {
		c.JSON(200, nil)
		return
	}
	c.JSON(200, map[string]any{
		"code": 0,
	})
}

func (h *helloHandler) Ping(c *gin.Context) {
	ctx, cancel := createContext()
	defer cancel()

	name := c.Query("name")
	id, err := h.hb.Ping(ctx, name)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	c.JSON(0, map[string]any{
		"message": "pong",
		"name":    name,
		"id":      id,
	})
}
