package jsoncoder

import (
	"io"
	"json"
	"os"
)

type Coder struct {
	enc *json.Encoder
	dec *json.Decoder
}

func NewCoder() *Coder {
	coder := &Coder{}

	return coder
}

func (coder *Coder) SetReadWriter(rw io.ReadWriter) {
	coder.enc = json.NewEncoder(rw)
	coder.dec = json.NewDecoder(rw)
}

func (coder *Coder) Encode(v interface{}) os.Error {
	return coder.enc.Encode(v)
}

func (coder *Coder) Decode(v interface{}) os.Error {
	return coder.dec.Decode(v)
}
