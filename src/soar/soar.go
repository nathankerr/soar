package soar

import (
	"net"
	"os"
)

type Consumer struct {
	laddr *net.UDPAddr //local addr
	raddr *net.UDPAddr //remote addr
	net string
}

func NewConsumer(laddr, raddr string) (*Consumer, os.Error) {
	consumer := new(Consumer)
	var err os.Error

	consumer.net = "udp"
	consumer.laddr, err = net.ResolveUDPAddr(consumer.net, laddr)
	if err != nil {
		return nil, err
	}

	consumer.raddr, err = net.ResolveUDPAddr(consumer.net, raddr)
	if err != nil {
		return nil, err
	}
	
	return consumer, nil
}

func (consumer *Consumer) Invoke (method string, args interface{}) (interface{}, os.Error) {
	c, err := net.DialUDP(consumer.net, consumer.laddr, consumer.raddr)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	_, err = c.WriteToUDP([]byte("ping"), consumer.raddr)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 1024)
	_, _, err = c.ReadFrom(buf)
	if err != nil {
		return nil, err
	}

	return string(buf), nil
}

type Server struct {
	addr *net.UDPAddr
	net string
}

func NewServer(addr string) (*Server, os.Error) {
	server := new(Server)
	var err os.Error

	server.net = "udp"
	server.addr, err = net.ResolveUDPAddr(server.net, addr)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (server *Server) Serve() os.Error {
	for {
		c, err := net.ListenUDP(server.net, server.addr)
		if err != nil {
			return err
		}
		defer c.Close()

		buf := make([]byte, 1024)
		n, addr, err := c.ReadFrom(buf)
		println(n, string(buf), "from", addr.String())

		c.WriteTo([]byte("pong"), addr)
	}

	return nil
}