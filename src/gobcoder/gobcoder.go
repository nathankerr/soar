package gobcoder

import (
	"encoding/gob"
	"io"
)

type Coder struct {
	enc *gob.Encoder
	dec *gob.Decoder
}

func NewCoder() *Coder {
	coder := &Coder{}

	return coder
}

func (coder *Coder) SetReadWriter(rw io.ReadWriter) {
	coder.enc = gob.NewEncoder(rw)
	coder.dec = gob.NewDecoder(rw)
}

func (coder *Coder) Encode(v interface{}) error {
	return coder.enc.Encode(v)
}

func (coder *Coder) Decode(v interface{}) error {
	return coder.dec.Decode(v)
}
