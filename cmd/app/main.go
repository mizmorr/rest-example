package main

import (
	"fmt"

	"github.com/mizmorr/rest-example/config"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	c := config.Get()
	fmt.Println(c)
	return nil
}
