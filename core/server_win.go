//go:build windows
// +build windows

package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
