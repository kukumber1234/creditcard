package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	stdin  = flag.Bool("stdin", false, "read")
	pick   = flag.Bool("pick", false, "choose")
	brands = flag.Bool("brands", false, "its brand")
)

func main() {
	var b []string
	var val bool = false
	var generat bool = false
	var info bool = false
	var issue bool = false
	b = append(b, os.Args...)
	if len(b) < 3 {
		os.Exit(1)
	}
	if b[1] == "issue" {
		issue = true
	}
	if b[1] == "information" {
		info = true
		if b[4] == "--stdin" {
			*stdin = true
			b = b[:4]
		}
	}
	if b[1] == "validate" {
		val = true
	}
	if b[2] == "--pick" {
		*pick = true
	}
	if b[2] == "--stdin" {
		*stdin = true
		b = b[:2]
	}
	if b[1] == "generate" {
		generat = true
	}
	if val && *stdin {
		number := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &number)
			if err != nil {
				break
			}
			b = append(b, number)
		}
		valid(b)
	} else if generat && *stdin {
		number := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &number)
			if err != nil {
				break
			}
			b = append(b, number)
		}
		gene(b)
	} else if info && *stdin {
		number := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &number)
			if err != nil {
				break
			}
			b = append(b, number)
		}
		informations(b)
	} else if val {
		valid(b)
	} else if generat {
		gene(b)
	} else if info {
		informations(b)
	} else if issue {
		issu(b)
	} else {
		fmt.Fprintln(os.Stderr, "Incorrect input")
		os.Exit(1)
	}
}