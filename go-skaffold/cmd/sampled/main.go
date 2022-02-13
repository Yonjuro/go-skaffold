package main

import (
	"github.com/yonjuro/go-skaffold/app"
)

func main() {
	server := app.NewServer()
	if err := server.Run(); err != nil {
		panic(err)
	}
}
