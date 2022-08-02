package main

import (
	"fmt"
	"github.com/HamidSajjadi/ushort/api"
	"github.com/HamidSajjadi/ushort/internal/config"
	"github.com/HamidSajjadi/ushort/internal/http-engine"
	"github.com/HamidSajjadi/ushort/internal/log"
	"github.com/HamidSajjadi/ushort/internal/repositories"
	"os"
	"os/signal"
	"syscall"
)

func initialize() {
	conf := config.New("ushort", "./config.yml")
	fmt.Printf("%+v", conf)
	logger := log.New(conf.LogLevel)
	urlRepo := repositories.NewInMemoryRepo()
	httpStub := http_engine.New(conf.Deployment, logger)
	handler := api.New(httpStub, urlRepo)
	handler.Run(conf.HttpAddress)
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
