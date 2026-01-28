package tcp

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	Address string
}

func NewServer(addr string) *Server {
	return &Server{Address: addr}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Server listening on", s.Address)

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		conn.Write([]byte(msg))
	}
}
