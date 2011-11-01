package soar

import (
	"fmt"
	"jsoncoder"
	"net"
	"os"
)

type Server struct {
	addr string
	listener net.Listener
	coder Coder
}

func NewServer(addr string) (*Server, os.Error) {
	coder := jsoncoder.NewCoder()
	return NewServerWithCoder(addr, coder)
}

func NewServerWithCoder(addr string, coder Coder) (server *Server, err os.Error) {
	server = &Server{ addr: addr,
		coder: coder,
	}

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
		server.coder.SetReadWriter(c)

		var invocation InvocationMessage
		server.coder.Decode(&invocation)
		fmt.Printf("%#v\n", invocation)

		server.coder.Encode("pong")

		c.Close()
	}

	return nil
}