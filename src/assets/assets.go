package main

import (
	"gobcoder"
	"log"
	"soar"
)

type Service int

func (s *Service) Ping(msg string) string {
	return "pong"
}

func main() {
	service := new(Service)
	coder := gobcoder.NewCoder()
	server, err := soar.NewServerWithCoder(":1234", service, coder)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Assets Server")
	server.Serve()
}