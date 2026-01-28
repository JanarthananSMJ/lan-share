package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Address string
}

func NewClient(addr string) *Client {
	return &Client{Address: addr}
}

func (c *Client) Start() error {
	conn, err := net.Dial("tcp", c.Address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Connected to server:", c.Address)

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("> ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			return err
		}

		response, err := serverReader.ReadString('\n')
		if err != nil {
			return err
		}

		fmt.Print("Server: ", response)
	}
}
