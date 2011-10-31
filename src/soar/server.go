package soar

import (
	"net"
	"os"
)

type Server struct {
	addr string
	listener net.Listener
}

func NewServer(addr string) (*Server, os.Error) {
	server := new(Server)
	var err os.Error

	server.addr = addr

	server.listener, err = net.Listen("tcp", server.addr)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (server *Server) Close() {
	server.listener.Close()
}

func (server *Server) Serve() os.Error {
	for {
		c, err := server.listener.Accept()
		if err != nil {
			return err
		}

		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		println(n, string(buf))

		c.Write([]byte("pong"))

		c.Close()
	}

	return nil
}