package main

import (
	"log"
	"soar"
)

func main() {
	server, err := soar.NewServer(":1234")
	if err != nil {
		panic(err)
	}

	log.Println("Starting Assets Server")
	server.Serve()
}