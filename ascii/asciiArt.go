package ascii

import (
	"fmt"
	"strings"
)

var color = map[string]string{
	"black":   "\033[30m",
	"red":     "\033[31m",
	"green":   "\033[32m",
	"yellow":  "\033[33m",
	"blue":    "\033[34m",
	"magenta": "\033[35m",
	"cyan":    "\033[36m",
	"white":   "\033[37m",
	"grey":    "\033[90m",
	"reset":   "\033[0m",
}

// Ascii prints the ASCII art based on the banner file and the input words.
// It takes two slices as input: fileArr and wordsArr.
func Ascii(fileArr []string, wordsArr []string, colorFlag string, letterFlag string) {
	colorFlag = strings.ToLower(colorFlag)
	letterToColor := make(map[rune]struct{})
	for _, v := range letterFlag {
		letterToColor[v] = struct{}{}
	}
	var count int

	for _, val := range wordsArr {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					if _, ok := letterToColor[v]; ok {
						fileArr[int(start)+i] = color[colorFlag] + fileArr[int(start)+i] + color["reset"]
						fmt.Print(fileArr[int(start)+i])
					} else {
						// Print the corresponding character of the banner file at the
						// calculated index
						fmt.Print(fileArr[int(start)+i])
					}
				}
				// Print a newline after each line of the ASCII character of the word
				fmt.Println()
			}
		} else {
			// If val is empty, print a new line
			count++
			if count < len(wordsArr) {
				fmt.Println()
			}
		}
	}
}
