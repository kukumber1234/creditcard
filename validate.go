package main

import "fmt"

func valid(b[] string) {
	for i := 2; i < len(b); i++ {
		if len(b[i]) < 13 {
			fmt.Println("INCORRECT")
		} else {
			if check(b[i]) {
				fmt.Println("OK")
			} else {
				fmt.Println("INCORRECT")
			}
		}
	}
}
