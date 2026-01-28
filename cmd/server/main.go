package main

import (
	"log"

	"github.com/JanarthananSMJ/lan-share/internal/tcp"
)

func main() {

	server := tcp.NewServer("0.0.0.0:9000")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
