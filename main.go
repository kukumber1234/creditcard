package main

import (
	"os"
)

func main() {
	var b []string
	b = append(b, os.Args...)
	if b[1] == "validate" {
		valid(b)
	}
}