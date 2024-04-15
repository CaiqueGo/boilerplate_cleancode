package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/boilerplate_cleancode/config"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Start(ctx)
	if err != nil {
		panic(err)
	}
	defer cfg.Close()
	server := gin.Default()
	userCreate := cfg.UserCreateHandlerHttp()
	server.POST("/user", userCreate.Create)

	server.Run(":8080")
}
