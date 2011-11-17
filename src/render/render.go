package main

import (
	"log"
	"soar"
)

type Render int

func (r *Render) Render(filename string) []byte {
	return []byte("hello from render! " + filename)
}

func main() {
	render := new(Render)
	server, err := soar.NewServer(":1233", render)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Render Server")
	server.Serve()
}
