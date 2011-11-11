package main

import (
	"log"
	"soar"
)

type Service int

func (s *Service) Echo(msg string) string {
	return msg + ", would you like fries with that?"
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
