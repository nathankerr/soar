package main

import (
	"exec"
	"path/filepath"
	"io/ioutil"
	"log"
	"os"
	"soar"
)

type Render int

func (r *Render) Render(filename string) []byte {
	assets_consumer, err := soar.NewConsumer(":1234")
	if err != nil {
		panic(err)
	}

	tmpdir, err := ioutil.TempDir("tmp/", "")
	if err != nil {
		panic(err)
	}

	returns, err := assets_consumer.Invoke("Get", filename)
	if err != nil {
		panic(err)
	}

	input, ok := returns[0].([]byte)
	if !ok {
		panic("could not recast assets return")
	}

	input_file := filepath.Join(tmpdir, "input.tex")
	ioutil.WriteFile(input_file, input, 0666)

	pdf_out := filepath.Join(tmpdir, "input.pdf")
	xelatex := exec.Command("xelatex",
		//"-o", pdf_out,
		"-output-directory=" + tmpdir,
		input_file,
	)
	xelatex.Stdout = os.Stdout
	xelatex.Stderr = os.Stderr
	if xelatex.Run() != nil {
		log.Fatal("pdfserver xelatex.Run:", err)
	}

	content, err := ioutil.ReadFile(pdf_out)
	if err != nil {
		log.Fatal("pdfserver read content:", err)
	}

	return content
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
