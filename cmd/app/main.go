package main

import (
	_ "github.com/mizmorr/rest-example/docs"
	app "github.com/mizmorr/rest-example/internal"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is example of RESTful api

//	@host		localhost:8080
//	@BasePath	/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Auth for user repo

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	return app.Run()
}
