package main

import (
	"io"
	"log"
	"net/http"
	"path"
	"soar"
)

func AssetsServer(w http.ResponseWriter, req *http.Request) {
	_, filename := path.Split(req.URL.Path)

	assets_client, err := soar.NewClient(":1234")
	if err != nil {
		panic(err)
	}

	if filename == "" {
		// List files
		io.WriteString(w, "<html><head><title>Asset List</title></head><body>")

		returns, err := assets_client.Invoke("List")
		if err != nil {
			panic(err)
		}

		files, ok := returns[0].([]string)
		if !ok {
			panic("cannot recast asset response")
		}

		for _, file := range files {
			io.WriteString(w, "<a href=\"/assets/"+file+"\">"+file+"</a> ")
			io.WriteString(w, "<a href=\"/render/"+file+"\">(as pdf)</a>")
			io.WriteString(w, "<br/>")
		}

		io.WriteString(w, "</body></html>")
	} else {
		// Show contents of a file
		returns, err := assets_client.Invoke("Get", filename)
		if err != nil {
			panic(err)
		}

		data, ok := returns[0].([]byte)
		if !ok {
			panic("cannot recast assets response")
		}

		w.Write(data)
	}
}

func RenderServer(w http.ResponseWriter, req *http.Request) {
	_, filename := path.Split(req.URL.Path)

	render_client, err := soar.NewClient(":1233")
	if err != nil {
		panic(err)
	}

	returns, err := render_client.Invoke("Render", filename)
	if err != nil {
		panic(err)
	}

	data, ok := returns[0].([]byte)
	if !ok {
		panic("cannot recast render response")
	}

	w.Write(data)
}

func main() {
	http.HandleFunc("/assets/", AssetsServer)
	http.HandleFunc("/render/", RenderServer)

	println("Starting Server on :12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
