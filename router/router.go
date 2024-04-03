package router

import (
	_ "api-sketch/handler"
	"github.com/dmzlingyin/utils/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	g := r.Group("/api")
	router.Register(g, "hello")
	return r
}
