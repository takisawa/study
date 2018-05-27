package main

import (
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	*c += WordCounter(len(strings.Fields(string(p))))

	return len(p), nil
}

func main() {
	var c WordCounter

	c.Write([]byte("Hello World"))
	fmt.Println(c)

	c = 0
	name := "Dolly"

	fmt.Fprintf(&c, "Hello %s how are you", name)
	fmt.Println(c)
}
