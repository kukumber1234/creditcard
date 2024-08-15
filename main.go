package main

import (
	"os"
	"flag"
	"fmt"
)

var (
	stdin = flag.Bool("stdin", false, "da")
)

func main() {
	var b []string
	var val bool = false
	var generat = false
	b = append(b, os.Args...)
	if len(b) < 3 {
		os.Exit(1)
	}
	if b[1] == "validate" {
		val = true
	}
	if b[2] == "--stdin" {
		*stdin = true
		b = b[:2]
	} else {
		os.Exit(1)
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
			b = append(b,number)
		}
		valid(b)
	} else if generat && *stdin{
		number := ""
		for {
			_, err := fmt.Fscan(os.Stdin, &number)
			if err != nil {
				break
			}
			b = append(b,number)
		}
		gene(b)
	} else if val {
		valid(b)
	} else if generat {
		gene(b)
	}
}






















// 	if b[1] == "validate" && b[2] == "--stdin" {
// 		*stdin = true
// 	} else if b[1] == "validate" {
// 		valid(b)
// 	} else if b[1] == "generate" {
		// gener(b)
// 	} else {
// 		fmt.Println("Incorrect imput")
// 		os.Exit(1)
// 	}
// 	if *stdin {
// 		number := ""
// 		for {
// 			_, err := fmt.Fscan(os.Stdin, &number)
// 			if err != nil {
// 				break
// 			}
// 			b = append(b,number)
// 		}
// 		valid(b)
// 	} 
// }
