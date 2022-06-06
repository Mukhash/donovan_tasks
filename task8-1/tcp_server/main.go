package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	port := flag.String("port", "8000", "a string")
	zone := flag.Int("zone", 0, "an int")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, *zone)
	}
}

func handleConn(c net.Conn, timezone int) {
	defer c.Close()

	for {
		_, err := io.WriteString(c, time.Now().Add(time.Hour*time.Duration(timezone)).Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}
}
