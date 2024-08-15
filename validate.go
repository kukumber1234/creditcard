package main

import (
	"fmt"
	"os"
)

func valid(b []string) {
	var incorr int = 0
	for i := 2; i < len(b); i++ {
		if len(b[i]) < 13 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			incorr++
		} else {
			if check(b[i]) {
				fmt.Println("OK")
			} else {
				fmt.Fprintln(os.Stderr, "INCORRECT")
				incorr++
			}
		}
	}
	if incorr >= 1 {
		os.Exit(1)
	}
}

func gene(b []string) {
	var inc int = 0
	for i := 2; i < len(b); i++ {
		if len(b[i]) < 13 {
			fmt.Fprintln(os.Stderr, "INCORRECT")
			inc++
		} else {
			proverit(b[i])
		}
	}
	if inc >= 1 {
		os.Exit(1)
	}
}
