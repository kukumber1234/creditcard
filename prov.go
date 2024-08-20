package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func check(b string) bool {
	var c []int
	var d int = 0
	var e int
	for i := 0; i < len(b); i++ {
		if rune(b[i]) >= 48 && rune(b[i]) <= 57 {
			c = append(c, int(b[i])-48)
			// fmt.Println(c)
		} else {
			return false
		}
	}
	if len(c)%2 != 0 {
		d = c[0]
		c = c[1:]
	}
	for i := 0; i < len(c); i++ {
		if i%2 == 0 {
			if c[i]*2 >= 10 {
				d += 1
				e = c[i] * 2 % 10
				d += e
			} else {
				d += c[i] * 2
			}
		} else {
			d += c[i]
		}
	}
	// fmt.Println(d)
	return d%10 == 0
}

func proverit(b string) {
	var f []int
	var g int = 0
	var kosu int = 0
	var kobeit int = 0
	for i := len(b) - 1; i >= 0; i-- {
		if b[len(b)-1] == '*' {
			if b[i] == '*' {
				g++
			}
		} else {
			os.Exit(1)
		}
	}

	b = b[:len(b)-g]
	for i := 0; i < len(b); i++ {
		if rune(b[i]) >= 48 && rune(b[i]) <= 57 {
			f = append(f, int(b[i])-48)
		} else {
			os.Exit(1)
		}
	}
	if g <= 4 {
		for i := len(f) - 1; i >= 0; i-- {
			kosu += f[i] * int(math.Pow(10, float64(kobeit)))
			kobeit++
		}
		kosu *= int(math.Pow(10, float64(g)))
		if *pick {
			for i := 0; i < int(math.Pow(10, float64(g))); i++ {
				random := rand.Intn(int(math.Pow(10, float64(g))))
				kerek := kosu + random
				if check(strconv.Itoa(kerek)) {
					fmt.Println(kerek)
					os.Exit(0)
				}
			}
		} else {
			for i := 0; i < int(math.Pow(10, float64(g))); i++ {
				kazirsan := kosu + i
				if check(strconv.Itoa(kazirsan)) {
					fmt.Println(kazirsan)
				}
			}
		}
	} else {
		os.Exit(1)
	}
}

func issue_generate(issueNumberAdd string) {
	var asteriskAdd int = 0
	var issueCountAdd []string
	var randomNumberString string
	var stringArray string
	var additionalStringArray string
	for i := 0; i < len(issueNumberAdd); i++ {
		if issueNumberAdd[i] == '*' {
			asteriskAdd++
		}
	}
	for i := len(issueNumberAdd)-1; i >= 0; i-- {
		if issueNumberAdd[i] != '*' {
			stringArray = issueNumberAdd[:i+1]
			additionalStringArray = issueNumberAdd[:i+1]
			break
		}
	}
	for i := 0; i < int(math.Pow(10, 3)); i++ {
		randomsan := rand.Intn(int(math.Pow(10, float64(asteriskAdd))))
		randomNumberString = strconv.Itoa(randomsan)
		issueCountAdd = append(issueCountAdd, randomNumberString)
	}
	for j := 0; j < len(issueCountAdd); j++ {
		stringArray += issueCountAdd[j]
		if check(stringArray) {
			fmt.Println(stringArray)
			os.Exit(0)
		} else {
			stringArray = additionalStringArray
		}
	}
}
