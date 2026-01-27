package main

import (
	"log"

	"github.com/JanarthananSMJ/lan-share/internal/tcp"
)

func main() {

	server := tcp.BuildServer("0.0.0.0:9000")
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
