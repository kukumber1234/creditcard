package helpe

import (
	"fmt"
	"os"
)

func CheckDir(dir string) {
	var countWord int
	var countNumber int
	var coutPeriods int

	if len(dir) < 3 || len(dir) > 63 {
		fmt.Fprintln(os.Stderr, "Names should be between 3 and 63 characters long (3 >= x >= 63)")
		os.Exit(1)
	}

	if dir[0] == '-' || dir[len(dir)-1] == '-' {
		fmt.Fprintln(os.Stderr, "Must not begin or end with a hyphen (-)")
		os.Exit(1)
	}

	for i := 0; i < len(dir); i++ {
		var checer bool = false
		if dir[i] <= '9' && dir[i] >= '0' {
			countNumber++
			checer = true
		}

		if dir[i] >= 97 && dir[i] <= 122 {
			countWord++
			checer = true
		}

		if dir[i] == 45 {
			if i < len(dir)-1 {
				if dir[i+1] == 45 {
					fmt.Fprintln(os.Stderr, "Must not contain two consecutive dashes (--)")
					os.Exit(1)
				}
			}
			checer = true
		}

		if dir[i] == 46 {

			coutPeriods++

			if i < len(dir)-1 {
				if dir[i+1] == 46 {
					fmt.Fprintln(os.Stderr, "Must not contain two consecutive periods (..)")
					os.Exit(1)
				}
			}
			checer = true
		}

		if dir[i] == 47 {
			checer = true
		}

		if !checer {
			fmt.Fprintln(os.Stderr, "Only lowercase letters (a-z), numbers (0-9), hyphnes (-), and dots (.) are allowwed")
			os.Exit(1)
		}
	}

	if countWord == 0 && countNumber != 0 && coutPeriods != 0 {
		fmt.Fprintln(os.Stderr, "Must not be formatted as an IP address (e.g. 192.168.0.1)")
		os.Exit(1)
	}
}
