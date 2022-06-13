package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Split(bufio.ScanRunes)
	for sc.Scan() {
		fmt.Print(sc.Text())
		fmt.Println()
		time.Sleep(2 * time.Second)
	}
}
