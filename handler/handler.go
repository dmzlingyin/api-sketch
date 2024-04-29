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
	return context.WithTimeout(context.Background(), 30*time.Second)
}

func handleError(c *gin.Context, err error) bool {
	if err != nil {
		log.Errorf("failed to handle, path: %s, err: %s", c.FullPath(), err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return true
	}
	return false
}

func response(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": data})
}
