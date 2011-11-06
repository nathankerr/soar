package main

import (
	"http"
	"io"
	"log"
	"soar"
)

func AssetsServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html><head><title>Asset List</title></head><body>")

	asset_consumer, err := soar.NewConsumer(":1234")
	if err != nil {
		panic(err)
	}

	returns, err := asset_consumer.Invoke("Ping", "first arg")
	if err != nil {
		panic(err)
	}

	response_string, ok := returns[0].(string)
	if !ok {
		panic("cannot recast response")
	}

	io.WriteString(w, response_string)
	io.WriteString(w, "</body></html>")
}

func main() {
	http.HandleFunc("/assets", AssetsServer)

	println("Starting Server on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}