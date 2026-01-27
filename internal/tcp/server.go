package tcp

import (
	"fmt"
	"net"
)

type Server struct {
	Address string
}

func BuildServer(address string) *Server {
	return &Server{Address: address}
}

func (s *Server) StartServer() error {
	listener, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Println("Server is Listening on : ", s.Address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Connection Established"))
}
