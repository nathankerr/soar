package soar

import (
	"os"
)

type Consumer struct {
	addr string
}

func NewConsumer(addr string) (*Consumer, os.Error) {
	consumer := new(Consumer)

	consumer.addr = addr

	return consumer, nil
}

func (consumer *Consumer) Invoke (method string, args interface{}) (interface{}, os.Error) {
	return "hello", nil
}