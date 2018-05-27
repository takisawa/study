package main

import (
	"fmt"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	*c += LineCounter(strings.Count(string(p), "\n"))

	return len(p), nil
}

func main() {
	var c LineCounter

	c.Write([]byte("Hello\nWorld\n"))
	fmt.Println(c)

	c = 0
	name := "Dolly"

	fmt.Fprintf(&c, "Hello %s\n how \nare \nyou", name)
	fmt.Println(c)
}
