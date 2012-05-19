package soar

import (
	"gobcoder"
	"net"
)

type Client struct {
	addr       string
	connection net.Conn
	coder      Coder
}

func NewClient(addr string) (*Client, error) {
	coder := gobcoder.NewCoder()
	return NewClientWithCoder(addr, coder)
}

func NewClientWithCoder(addr string, coder Coder) (client *Client, err error) {
	client = &Client{addr: addr,
		coder: coder,
	}

	client.connection, err = net.Dial("tcp", client.addr)
	if err != nil {
		return nil, err
	}

	client.coder.SetReadWriter(client.connection)

	return client, nil
}

func (client *Client) Close() {
	client.connection.Close()
}

func (client *Client) Invoke(capability string, args ...interface{}) (returns []interface{}, err error) {
	request := &Request{Capability: capability,
		Args: args,
	}
	err = client.coder.Encode(request)
	if err != nil {
		return nil, err
	}

	var response Response
	err = client.coder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Returns, response.Err
}
