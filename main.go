package main

import (
	"fmt"
	"github.com/HamidSajjadi/ushort/api"
	"github.com/HamidSajjadi/ushort/internal/config"
	"github.com/HamidSajjadi/ushort/internal/db"
	"github.com/HamidSajjadi/ushort/internal/http-engine"
	"github.com/HamidSajjadi/ushort/internal/log"
	"github.com/HamidSajjadi/ushort/internal/repositories"
	"os"
	"os/signal"
	"syscall"
)

func initialize() {
	conf := config.New("ushort", "./config.yml")
	logger := log.New(conf.LogLevel)
	urlRepo := chooseRepository(conf)
	httpStub := http_engine.New(conf.Deployment, logger)
	handler := api.New(httpStub, urlRepo)
	handler.Run(conf.HttpAddress)
}

func chooseRepository(conf *config.Config) repositories.URLRepository {
	switch conf.DatabaseType {

	case config.DatabaseTypeRedis:
		redisDB, err := db.NewRedis(conf.Redis.Address, conf.Redis.Password)
		if err != nil {
			panic(fmt.Sprintf("could not connect to redis at `%s`, error: %v", conf.Redis.Address, err))
		}
		return repositories.NewRedisRepo(redisDB)
	case config.DatabaseTypeInMemory:
		return repositories.NewInMemoryRepo()
	default:
		panic(fmt.Sprintf("repository type `%s` not implemented", conf.DatabaseType))
	}
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
