package main

import (
	"bufio"
	"fmt"
	"os"
)

func informations(b []string) {
	var brandsname string
	var issuersname string
	var minus string
	var minus1 string
	brandOrIssuer := b[2]
	brandOrIssuer1 := b[3]
	for i := 0; i < len(b[2]); i++ {
		if brandOrIssuer[i] == '=' {
			brandsname = brandOrIssuer[i+1:]
			minus = brandOrIssuer[:i+1]
		}
	}
	for i := 0; i < len(b[3]); i++ {
		if brandOrIssuer1[i] == '=' {
			issuersname = brandOrIssuer1[i+1:]
			minus1 = brandOrIssuer1[:i+1]
		}
	}
	if (minus == "--brands=" && minus1 == "--issuers=") || (minus1 == "--issuers=" && minus == "--brands=") {
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
			if *cardbrandnotsan == len(gg) {
				fmt.Fprintln(os.Stderr, "Card Brand: -")
			}
			if *issuernotsan == len(jj) {
				fmt.Fprintln(os.Stderr, "Card Issuer: -")
			}
		} else {
			fmt.Fprintln(os.Stderr, "Correct: no")
			fmt.Fprintln(os.Stderr, "Card Brand: -")
			fmt.Fprintln(os.Stderr, "Card Issuer: -")
			os.Exit(1)
		}
	} else {
		fmt.Fprintln(os.Stderr, "Incorrect input")
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
			} else {
				*cardbrandnotsan++
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
			} else {
				*issuernotsan++
			}
		}
	}
}
