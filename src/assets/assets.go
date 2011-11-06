package main

import (
	"log"
	"soar"
)

type Service int

func (s *Service) Ping(msg string) string {
	return "pong"
}

func main() {
	service := new(Service)
	server, err := soar.NewServer(":1234", service)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Assets Server")
	server.Serve()
}