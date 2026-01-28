package main

import (
	"flag"
	"log"

	"github.com/JanarthananSMJ/lan-share/internal/tcp"
)

func main() {
	IpFlag := flag.String("ip", "", "Use Server IP to connect")
	flag.Parse()
	client := tcp.NewClient(*IpFlag)
	if err := client.Start(); err != nil {
		log.Fatal(err)
	}

}
