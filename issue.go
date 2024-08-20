package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func issu(b []string) {
	var issuesankosu string
	var issuesan string = ""
	var brandneshe int = 0
	var issuerneshe int = 0
	var issuerchtosandar string
	var brandchtosandar string
	var brandchtosandar1 string
	var issuerchtosandar1 string
	var brandsan []string
	var issuersan string
	var brandname string
	var issuername string
	var brandname1 string
	var issuername1 string
	var brandcost int = 0
	var issuecost int = 0
	var btxt string
	var itxt string
	var minus2 string
	var minus3 string
	brandOrIssuer2 := b[2]
	brandOrIssuer3 := b[3]
	for i := 0; i < len(b[2]); i++ {
		if brandOrIssuer2[i] == '=' {
			btxt = brandOrIssuer2[i+1:]
			minus2 = brandOrIssuer2[:i+1]
		}
	}
	for i := 0; i < len(b[3]); i++ {
		if brandOrIssuer3[i] == '=' {
			itxt = brandOrIssuer3[i+1:]
			minus3 = brandOrIssuer3[:i+1]
		}
	}
	if (len(b) >= 2 && minus2 == "--brands=" && minus3 == "--issuers=") || (len(b) >= 2 && minus3 == "--brands=" && minus2 == "--issuers=") {
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
			fmt.Fprintln(os.Stderr, "Missing '=' in", brandname)
			os.Exit(1)
		}
		if issuecost == len(issuername) {
			fmt.Fprintln(os.Stderr, "Missing '=' in", issuername)
			os.Exit(1)
		}
		brandchto := ChitatFile2(btxt)
		issuerschto := ChitatFile2(itxt)
		for i := 0; i < len(brandchto); i++ {
			brandchtosandar = brandchto[i]
			for j := 0; j < len(brandchtosandar); j++ {
				if brandchtosandar[j] == ':' {
					brandchtosandar1 = brandchtosandar[:j]
					if brandchtosandar1 == brandname1 {
						brandsan = append(brandsan, brandchtosandar[j+1:])
						brandneshe++
					}
				}
			}
		}
		for i := 0; i < len(issuerschto); i++ {
			issuerchtosandar = issuerschto[i]
			for j := 0; j < len(issuerchtosandar); j++ {
				if issuerchtosandar[j] == ':' {
					issuerchtosandar1 = issuerchtosandar[:j]
					if issuerchtosandar1 == issuername1 {
						issuersan = issuerchtosandar[j+1:]
						issuerneshe++
					}
				}
			}
		}
		if brandneshe <= 0 {
			fmt.Fprintln(os.Stderr, "No such brand in brands.txt")
			os.Exit(1)
		}
		if issuerneshe <= 0 {
			fmt.Fprintln(os.Stderr, "No such issuer in issuers.txt")
			os.Exit(1)
		}
		for i := 0; i < len(brandsan); i++ {
			if brandsan[i] == issuersan[:len(brandsan[i])] {
				issuesan += issuersan
			}
		}
		if len(issuesan) == 0 {
			fmt.Fprintln(os.Stderr, "They are not suitable for each other")
			os.Exit(1)
		} else {
			issuesankosu += issuesan
			if brandname1 == "MASTERCARD" || brandname1 == "DISCOVER" || brandname1 == "JCB" || brandname1 == "InstaPayment" {
				for i := 0; i < 16-len(issuesan); i++ {
					issuesankosu += "*"
				}
			} else if brandname1 == "AMEX" {
				for i := 0; i < 15-len(issuesan); i++ {
					issuesankosu += "*"
				}
			} else if brandname1 == "DinerdClubCarteBlanche" || brandname1 == "DinerdClubInternational" {
				for i := 0; i < 14-len(issuesan); i++ {
					issuesankosu += "*"
				}
			} else if brandname1 == "VISA" {
				min := 13
				max := 16
				Visa := rand.Intn(max-min+1) + min
				for i := 0; i < Visa-len(issuesan); i++ {
					issuesankosu += "*"
				}
			}
			issue_generate(issuesankosu)
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
