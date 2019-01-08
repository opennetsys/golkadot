package main

import (
	"github.com/c3systems/go-substrate/client"
	"github.com/c3systems/go-substrate/logger"
)

func main() {
	// TODO: implement
	config := &client.Config{}
	cl := client.NewClient()
	cl.Start(config)

	logger.Set(logger.ContextHook{}, false)
}
