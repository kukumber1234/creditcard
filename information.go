package main

import (
	"bufio"
	"fmt"
	"os"
)

func informations(b []string) {
	var brandsname string = "brands.txt"
	var issuersname string = "issuers.txt"
	if b[2] == "--brands=brands.txt" && b[3] == "--issuers=issuers.txt" {
		*brands = true
	}
	if *brands {
		san := b[4]
		fmt.Println(san)
		if check(san) {
			fmt.Println("Correct: yes")
			gg := ChitatFile(brandsname)
			jj := ChitatFile(issuersname)
			for i := 0; i < len(gg); i++ {
				brandfound(gg[i], san)
			}
			for i := 0; i < len(jj); i++ {
				issuerfound(jj[i], san)
			}
		} else {
			fmt.Println("Correct: no")
			fmt.Println("Card Brand: -")
			fmt.Println("Card Issuer: -")
		}
	} else {
		fmt.Println("Incorrect input")
		os.Exit(1)
	}
}

func ChitatFile(filename string) []string {
	f, _ := os.Open(filename)
	defer f.Close()
	s := bufio.NewScanner(f)
	d := []string{}
	for s.Scan() {
		d = append(d, s.Text())
	}
	return d
}

func brandfound(gg string, san string) {
	var sohr string
	for i := 0; i < len(gg); i++ {
		if gg[i] == ':' {
			sohr = gg[i+1:]
			if sohr == san[:len(sohr)] {
				fmt.Println("Card Brand:", gg[:len(gg)-len(sohr)-1])
			}
		}
	}
}

func issuerfound(jj string, san string) {
	var sohro string
	for i := 0; i < len(jj); i++ {
		if jj[i] == ':' {
			sohro = jj[i+1:]
			if sohro == san[:len(sohro)] {
				fmt.Println("Card Issuer:", jj[:len(jj)-len(sohro)-1])
			}
		}
	}
}