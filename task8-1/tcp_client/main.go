package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	type clockSet struct {
		zone   string
		source *bufio.Reader
	}

	var clocks []clockSet

	flag.Parse()

	colWidth := 10
	for _, arg := range flag.Args() {
		argParts := strings.Split(arg, "=")
		if len(argParts) == 2 {

			conn, err := net.Dial("tcp", argParts[1])
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()

			clocks = append(clocks, clockSet{
				zone:   argParts[0],
				source: bufio.NewReader(conn),
			})

			if len(argParts[0]) > colWidth {
				colWidth = len(argParts[0])
			}
		}
	}

	fmt.Print("\x1B[0;0f\x1B[2J")
	for i, clock := range clocks {
		fmt.Printf("\x1B[1;%df%s", i*colWidth, clock.zone)
	}
	fmt.Println()

	for {
		for i, clock := range clocks {
			fmt.Printf("\x1B[2;%df", i*colWidth)
			line, err := clock.source.ReadString('\n')
			if err != nil {
				fmt.Printf("\x1B[2;%dfError   ", i*colWidth)
				continue
			}
			fmt.Printf("\x1B[2;%df%s", i*colWidth, line)
		}
	}
}
