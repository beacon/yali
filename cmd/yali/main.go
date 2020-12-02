package main

import (
	"log"

	"github.com/beacon/yali/cmd/yali/app"
)

func main() {
	command := app.NewYaliCommand()
	if err := command.Execute(); err != nil {
		log.Fatalln("Failed to execute:", err)
	}
}
