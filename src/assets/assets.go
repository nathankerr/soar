package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"soar"
)

type Assets struct {
	path string
}

func NewAssets() *Assets {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(path, "assets")

	return &Assets{path: path}
}

func (a *Assets) List() []string {
	dir, err := os.Open(a.path)
	if err != nil {
		panic(err)
	}

	names, err := dir.Readdirnames(-1) // -1 means to read all names
	if err != nil {
		panic(err)
	}

	return names
}

func (a *Assets) Get(filename string) []byte {
	path := filepath.Join(a.path, filename)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

func main() {
	assets := NewAssets()
	server, err := soar.NewServer(":1234", assets)
	if err != nil {
		panic(err)
	}

	log.Println("Starting Assets Server")
	server.Serve()
}
