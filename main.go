package main

import (
	"api-sketch/router"
	"github.com/dmzlingyin/utils/log"
	"net/http"
	"time"
)

func main() {
	handler := router.Router()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   5 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Errorf("listen failed: %v", err)
	}
}
