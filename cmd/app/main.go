package main

import (
	"context"

	store "github.com/mizmorr/rest-example/store"
)

func main() {

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()
	_, err := store.New(ctx)
	if err != nil {
		return err
	}
	return nil
}
