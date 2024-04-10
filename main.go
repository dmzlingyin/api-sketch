package main

import (
	"api-sketch/router"
	"errors"
	"fmt"
	"github.com/dmzlingyin/utils/config"
	"github.com/dmzlingyin/utils/log"
	"net/http"
	"os"
	"time"
)

func main() {
	var p string
	if p = os.Getenv("PROFILE"); p != "" {
		config.SetProfile(fmt.Sprintf("config/%s.json", p))
	}

	port := config.Get("app.port").String()
	s := &http.Server{
		Addr:           port,
		Handler:        router.Router(p),
		ReadTimeout:    5 * time.Minute,
		WriteTimeout:   5 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infof("server listening on %s", port)

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Errorf("server exited!!! err: %s", err)
	}
}
