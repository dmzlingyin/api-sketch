package main

import (
	_ "api-sketch/handler"
	"api-sketch/middleware"
	"fmt"
	"github.com/dmzlingyin/utils/config"
	"github.com/dmzlingyin/utils/router"
	"github.com/dmzlingyin/utils/server"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	var p string
	if p = os.Getenv("PROFILE"); p != "" {
		config.SetProfile(fmt.Sprintf("config/%s.json", p))
	}
	server.Serve(config.GetString("app.port"), handler(p))
}

func handler(profile string) *gin.Engine {
	if profile == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(router.DefaultCors())
	r.Use(middleware.Auth())

	g := r.Group("/api")
	register(g)
	return r
}

func register(g *gin.RouterGroup) {
	const prefix = "handler."
	router.Register(g, prefix, "hello")
}
