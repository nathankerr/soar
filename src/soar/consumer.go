package soar

import (
	"gobcoder"
	"net"
	"os"
)

type Consumer struct {
	addr string
	connection net.Conn
	coder Coder
}

func NewConsumer(addr string) (*Consumer, os.Error) {
	coder := gobcoder.NewCoder()
	return NewConsumerWithCoder(addr, coder)
}

func NewConsumerWithCoder(addr string, coder Coder) (consumer *Consumer, err os.Error) {
	consumer = &Consumer{ addr: addr,
		coder: coder,
	}

	consumer.connection, err = net.Dial("tcp", consumer.addr)
	if err != nil {
		return nil, err
	}

	consumer.coder.SetReadWriter(consumer.connection)
	
	return consumer, nil
}

func (consumer *Consumer) Close() {
	consumer.connection.Close()
}

func (consumer *Consumer) Invoke (capability string, args ...interface{}) (returns []interface{}, err os.Error) {
	request := &Request{ Capability: capability,
		Args: args,
	}
	err = consumer.coder.Encode(request)
	if err != nil {
		return nil, err
	}

	var response Response
	err = consumer.coder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Returns, response.Err
}