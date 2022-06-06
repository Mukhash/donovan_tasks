package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(s []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(s)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil { //asdfdsd
		return int(*c), err
	}
	return int(*c), nil
}

func (c *LineCounter) Write(s []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(s)))
	for scanner.Scan() {
		*c++
	}
	if err := scanner.Err(); err != nil {
		return int(*c), err
	}
	return int(*c), nil
}

func main() {
	words := "one two three four five"
	lines := "one\ntwo\nthree\nfour\nfive\n"

	var wc WordCounter = 0
	var lc LineCounter = 0

	wc.Write([]byte(words))
	lc.Write([]byte(lines))

	fmt.Printf("num of words: %d\n", wc)
	fmt.Printf("num of lines: %d\n", lc)
}
