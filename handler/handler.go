package handler

import (
	"context"
	"github.com/dmzlingyin/utils/ioc"
	"time"
)

func init() {
	ioc.Put(NewHelloHandler, "handler.hello")
}

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Minute)
}
