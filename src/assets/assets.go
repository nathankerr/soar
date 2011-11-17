package main

import (
	"path/filepath"
	"log"
	"soar"
	"os"
)

type Assets int

func (s *Assets) List() []string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(path, "assets")

	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	names, err := dir.Readdirnames(-1) // -1 means to read all names
	if err != nil {
		panic(err)
	}

	return names
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
