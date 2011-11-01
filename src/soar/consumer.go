package soar

import (
	"jsoncoder"
	"net"
	"os"
)

type Consumer struct {
	addr string
	connection net.Conn
	coder Coder
}

func NewConsumer(addr string) (*Consumer, os.Error) {
	coder := jsoncoder.NewCoder()
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

func (consumer *Consumer) Invoke (capability string, args ...interface{}) (interface{}, os.Error) {
	invocation := &InvocationMessage{ Capability: capability,
		Args: args,
	}
	consumer.coder.Encode(invocation)

	var msg string
	consumer.coder.Decode(&msg)

	return msg, nil
}