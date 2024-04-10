package handler

import (
	"context"
	"github.com/dmzlingyin/utils/ioc"
	"github.com/dmzlingyin/utils/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	ioc.Put(NewHelloHandler, "handler.hello")
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Minute)
}

func response(c *gin.Context, err error, data any) {
	if err != nil {
		log.Errorf("failed to handle, path: %s, err: %s", c.FullPath(), err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}
