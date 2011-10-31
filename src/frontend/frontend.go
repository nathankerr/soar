package main

import (
	"http"
	"io"
	"log"
	"soar"
)

func process() {

	// consumer, _ := soar.NewConsumer(":1234")

	// soar defaults to request/response style call
	// if something else is used, then extra work should be needed to use it
	// for example, if coordination is used, then: consumer.InteractionMode = COORDINATION and change to the coordination API
	// Request/Response in a coordinated system simply sets the point the response should be sent to to this client and not another service. consumer.InvokeCoordinated("method", args, input, destination)
	// A transactional system would wrap a single invocation in a transaction. transaction := consumer.NewTransaction(); transaction.Invoke("method", args); transaction.Commit(); transaction.Abort()
	// to handle Async and Sync methods, the synchrous method could always wrap the async method. This would allow synchronous invocation on an asynchronous system.
	// The idea is to maintain the API so that services and their consumers do not have to be updated when a new feature is added.
	// Phase 1: write synchronous, request/reply; use to implement a "Complete SOA"
	// args := Maths{1, 2}
	// response, _ := consumer.Invoke("multiply", args) //response is a interface{}
	// mult_resp, err := response.(MathsResponse)
}

func RootServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<html><head><title>Asset List</title></head><body>")

	asset_consumer, err := soar.NewConsumer(":1233", ":1234")
	if err != nil {
		panic(err)
	}

	response, err := asset_consumer.Invoke("method", nil)
	if err != nil {
		panic(err)
	}

	response_string, ok := response.(string)
	if !ok {
		panic("cannot recast response")
	}

	io.WriteString(w, response_string)
}

func main() {
	http.HandleFunc("/", RootServer)

	println("Starting Server on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.String())
	}
}