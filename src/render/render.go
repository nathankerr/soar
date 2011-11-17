package main

import (
	"log"
	"soar"
)

type Render int

func (r *Render) Render(filename string) []byte {
	assets_consumer, err := soar.NewConsumer(":1234")
	if err != nil {
		panic(err)
	}

	returns, err := assets_consumer.Invoke("Get", filename)
	if err != nil {
		panic(err)
	}

	data, ok := returns[0].([]byte)
	if !ok {
		panic("could not recast assets return")
	}

	return data
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
