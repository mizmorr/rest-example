package main

import (

	"github.com/mizmorr/rest-example/internal"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	return internal.Run()
}
