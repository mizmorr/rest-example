package main

import (
	app "github.com/mizmorr/rest-example/internal"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	return app.Run()
}
