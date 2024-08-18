package main

import (
	"fmt"
	"os"
	"bufio"
)

func issu(b []string) {
	var brandname string
	var issuername string
	var brandname1 string
	var issuername1 string
	var brandcost int = 0
	var issuecost int = 0
	btxt := "brands.txt"
	// itxt := "issuers.txt"
	if b[2] == "--brands=brands.txt" && b[3] == "--issuers=issuers.txt" {
		*brands = true
		b = b[4:]
	} else {
		fmt.Fprintln(os.Stderr, "Incorrect input")
		os.Exit(1)
	}
	if *brands {
		brandname = b[0]
		issuername = b[1]
		for i := 0; i < len(brandname); i++ {
			if brandname[i] == '=' {
				if brandname[:i+1] != "--brand=" {
					fmt.Fprintln(os.Stderr, "Incorrect input")
					os.Exit(1)
				}
				brandname1 = brandname[i+1:]
			} else {
				brandcost++
			}
		}
		for i := 0; i < len(issuername); i++ {
			if issuername[i] == '=' {
				if issuername[:i+1] != "--issuer=" {
					fmt.Fprintln(os.Stderr, "Incorrect input")
					os.Exit(1)
				}
				issuername1 = issuername[i+1:]
			} else {
				issuecost++
			}
		}
		if brandcost == len(brandname) {
			fmt.Fprintln(os.Stderr, "Missing '=' in",brandname)
			os.Exit(1)
		}
		if issuecost == len(issuername) {
			fmt.Fprintln(os.Stderr, "Missing '=' in",issuername)
			os.Exit(1)
		}
		fmt.Println(brandname1)
		fmt.Println(issuername1)
		brandchto := ChitatFile2(btxt)
		// issuerschto := ChitatFile2(itxt)
		for i := 0; i < len(brandchto); i++ {
			if brandname == brandchto[i] {
				fmt.Println("DA")
			} 
		}
	}
}


func ChitatFile2(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()
	s := bufio.NewScanner(f)
	d := []string{}
	for s.Scan() {
		d = append(d, s.Text())
	}
	return d
}