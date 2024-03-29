package soar

import (
	"gobcoder"
	"net"
	"reflect"
)

type Server struct {
	addr     string
	listener net.Listener
	coder    Coder
	service  interface{} // variable representing the service and its capabilities
}

func NewServer(addr string, service interface{}) (*Server, error) {
	coder := gobcoder.NewCoder()
	return NewServerWithCoder(addr, service, coder)
}

func NewServerWithCoder(addr string, service interface{}, coder Coder) (server *Server, err error) {
	server = &Server{addr: addr,
		service: service,
		coder:   coder,
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

func (server *Server) Serve() error {
	for {
		c, err := server.listener.Accept()
		if err != nil {
			return err
		}
		server.coder.SetReadWriter(c)

		var request Request
		server.coder.Decode(&request)

		response := server.call(request)

		server.coder.Encode(response)

		c.Close()
	}

	return nil
}

func (server *Server) call(request Request) *Response {
	service := reflect.ValueOf(server.service)
	capability := service.MethodByName(request.Capability)
	if !capability.IsValid() {
		panic("Capability not found")
	}

	args := make([]reflect.Value, len(request.Args))
	for k, v := range request.Args {
		args[k] = reflect.ValueOf(v)
	}

	returns := capability.Call(args)

	response := &Response{Returns: make([]interface{}, len(returns))}
	for k, v := range returns {
		response.Returns[k] = v.Interface()
	}

	return response
}
