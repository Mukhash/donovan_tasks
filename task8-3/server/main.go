package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	fmt.Println("started listening on 8000...")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("started accepting conns ...")
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, shout)
	time.Sleep(delay)
	fmt.Fprintln(c, strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}

	c.Close()
	fmt.Println("conn closed...")
}
