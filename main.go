package main

import (
	"github.com/c3systems/go-substrate/log"
)

func main() {
	log.Set(log.ContextHook{}, false)
}
