package ascii

import (
	"fmt"
	"strings"
)

func Ascii(fileArr []string, wordsArr []string, option []string, color string) {
	var count int
	reset := "\033[0m"

	opt := strings.Join(option, "")
	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					if len(opt) == 0 {
						start := (v - 32) * 9
						fmt.Print(color + fileArr[int(start)+i] + reset)
					} else if strings.Contains(opt, string(v)) {
						start := (v - 32) * 9
						fmt.Print(color + fileArr[int(start)+i] + reset)
					} else {
						start := (v - 32) * 9
						fmt.Print(fileArr[int(start)+i])
					}
				}
				fmt.Println()
			}
		} else {
			count++
			if count < len(wordsArr) {
				fmt.Println()
			}
		}
	}
}
