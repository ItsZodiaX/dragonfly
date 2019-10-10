package main

import (
	"github.com/dragonfly-tech/dragonfly/dragonfly"
)

func main() {
	server := dragonfly.New(nil, nil)
	if err := server.Run(); err != nil {
		panic(err)
	}
	for {
		server.Accept()
	}
}
