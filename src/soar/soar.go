package soar

import (
	"io"
)

type Coder interface {
	SetReadWriter(rw io.ReadWriter)
	Encode(v interface{}) error
	Decode(v interface{}) error
}

type Request struct {
	Capability string
	Args       []interface{}
}

type Response struct {
	Err     error
	Returns []interface{}
}
