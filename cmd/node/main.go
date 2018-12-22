package main

import "github.com/c3systems/go-substrate/logger"

func main() {
	logger.Set(logger.ContextHook{}, false)
}
