package internal

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/rest-example/internal/controller"
	"github.com/mizmorr/rest-example/internal/server"
	"github.com/mizmorr/rest-example/pkg/logger"
	"github.com/mizmorr/rest-example/service"
	"github.com/mizmorr/rest-example/store"
	"github.com/pkg/errors"
)

func Run() error {
	ctx := context.Background()


	//logger
	l:= logger.Get()

	store, err:= store.New(ctx)
	if err!= nil {
		return errors.Wrap(err,"store.New failed")
	}

	svc,err := service.NewUserWebService(store,ctx)

	if err!= nil {
		return errors.Wrap(err, "service.NewUserWebService failed")
	}

	//user controller

	userController:= controller.NewUsers(ctx,svc,l)

	handler:= gin.New()

	server.NewRouter(handler,userController)
	l.Info().Msg("Server is listening ...")
	handler.Run(":8080")
	return nil
}
