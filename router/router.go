package router

import (
	_ "api-sketch/handler"
	"fmt"
	"github.com/dmzlingyin/utils/ioc"
	"github.com/dmzlingyin/utils/log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

func Router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	g := r.Group("/api")
	register(g, "hello")
	return r
}

func register(g *gin.RouterGroup, name string) {
	ins, err := ioc.TryFind(fmt.Sprintf("handler.%s", name))
	if err != nil {
		panic(err)
	}
	v := reflect.Indirect(reflect.ValueOf(ins))
	if v.Kind() != reflect.Struct {
		panic("invalid handler type: " + name)
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		h, ok := f.Interface().(gin.HandlerFunc)
		if !ok {
			continue
		}
		field := t.Field(i)
		path := field.Tag.Get("path")
		relativePath := fmt.Sprintf("/%s%s", name, path)
		method := field.Tag.Get("method")
		registerMethod(g, method, relativePath, h)
	}
}

func registerMethod(g *gin.RouterGroup, method, rPath string, h gin.HandlerFunc) {
	switch strings.ToLower(method) {
	case "get":
		g.GET(rPath, h)
	case "post":
		g.POST(rPath, h)
	case "put":
		g.PUT(rPath, h)
	case "delete":
		g.DELETE(rPath, h)
	case "patch":
		g.PATCH(rPath, h)
	default:
		panic("invalid method: " + method)
	}
	log.Infof("handler registered: [%s], path: %s", method, rPath)
}
