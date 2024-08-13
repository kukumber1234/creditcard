package main

func check(b string) bool {
	var c[] int
	var d int = 0
	var e int
	for i := 0; i < len(b); i++ {
		if rune(b[i]) >= 48 && rune(b[i]) <= 57 {
			c = append(c,int(b[i]) - 48)
			// fmt.Println(c)
		} else {
			return false
		}
	}
	for i := 0; i < len(c); i++ {
		if i % 2 == 0 {
			if c[i] * 2 > 10 {
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
	if d % 10 == 0 {
		return true
	} else {
		return false
	}
}