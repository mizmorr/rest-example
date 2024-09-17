package internal

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/rest-example/config"
	"github.com/mizmorr/rest-example/internal/controller"
	"github.com/mizmorr/rest-example/internal/router"
	"github.com/mizmorr/rest-example/pkg/logger"
	"github.com/mizmorr/rest-example/pkg/server"
	"github.com/mizmorr/rest-example/service"
	"github.com/mizmorr/rest-example/store"
	"github.com/mizmorr/rest-example/store/redis"

	"github.com/pkg/errors"
)

func Run() error {
	ctx := context.Background()

	cfg := config.Get()
	//logger
	l := logger.Get()

	store, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	svc, err := service.NewUserWebService(store, ctx)

	if err != nil {
		return errors.Wrap(err, "service.NewUserWebService failed")
	}

	cache := redis.New(0, "localhost:6379", "1234")

	if err = cache.Setup(ctx); err != nil {
		return errors.Wrap(err, "cache.Setup failed")
	}
	cacheController := controller.NewCache(cache)

	//user controller

	userController := controller.NewUsers(ctx, svc)

	handler := gin.New()

	router.NewRouter(handler, userController, cacheController)

	httpServer := server.New(handler)
	l.Info().Msg(fmt.Sprintf("Server is running on %v ...", cfg.HTTPAddress))
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msg("[app.Run] - signal " + s.String())
	case err = <-httpServer.Notify():
		l.Error().Err(fmt.Errorf("[app.Run] - httpServer.Notify " + err.Error()))
	}

	return httpServer.Shutdown()

}
