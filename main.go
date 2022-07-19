package main

import (
	"fmt"
	"github.com/HamidSajjadi/ushort/api"
	"github.com/HamidSajjadi/ushort/internal/repositories"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func initialize() {
	urlRepo := repositories.NewInMemoryRepo()
	httpStub := gin.Default()

	handler := api.New(httpStub, urlRepo)
	handler.Run("localhost:9090")
}
func wait() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		fmt.Println("Signal received: ", sig)
		done <- true
	}()
	<-done
}

func main() {
	initialize()
	wait()
}
