package ascii

import (
	"fmt"
	"strings"
)

func Ascii(fileArr []string, wordsArr []string, letterToColor string, color string) {
	var count int
	reset := "\033[0m"

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if len(opt) == 0 {
						fmt.Print(color + fileArr[int(start)+i] + reset)
					} else if strings.Contains(opt, string(v)) {
						fmt.Print(color + fileArr[int(start)+i] + reset)
					} else {
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
