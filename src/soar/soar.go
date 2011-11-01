package soar

import (
	"io"
	"os"
)

type Coder interface {
	SetReadWriter(rw io.ReadWriter)
	Encode(v interface{}) os.Error
	Decode(v interface{}) os.Error
}

type InvocationMessage struct {
	Capability string
	Args []interface{}
}