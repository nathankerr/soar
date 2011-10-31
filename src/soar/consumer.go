package soar

import (
	"net"
	"os"
)

type Consumer struct {
	addr string
	connection net.Conn
}

func NewConsumer(addr string) (consumer *Consumer, err os.Error) {
	consumer = &Consumer{ addr: addr }

	consumer.connection, err = net.Dial("tcp", consumer.addr)
	if err != nil {
		return nil, err
	}
	
	return consumer, nil
}

func (consumer *Consumer) Close() {
	consumer.connection.Close()
}

func (consumer *Consumer) Invoke (method string, args interface{}) (interface{}, os.Error) {

	_, err := consumer.connection.Write([]byte("ping"))
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	_, err = consumer.connection.Read(buf)
	if err != nil {
		return nil, err
	}

	return string(buf), nil
}