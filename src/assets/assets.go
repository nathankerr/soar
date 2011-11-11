package main

import (
	"log"
	"soar"
)

type Assets int

func (s *Assets) Echo(msg string) string {
	return msg + ", would you like fries with that?"
}

func main() {
	assets := new(Assets)
	server, err := soar.NewServer(":1234", assets)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Assets Server")
	server.Serve()
}
