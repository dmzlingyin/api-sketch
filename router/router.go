package router

import (
	_ "api-sketch/handler"
	"github.com/dmzlingyin/utils/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(profile string) *gin.Engine {
	if profile == "prd" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	g := r.Group("/api")
	register(g)
	return r
}

func register(g *gin.RouterGroup) {
	const prefix = "handler."
	router.Register(g, prefix, "hello")
}
