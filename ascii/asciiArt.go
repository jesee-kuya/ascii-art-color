package ascii

import (
	"fmt"
	"strings"
)

var color = map[string]string{
	"black":       "\033[30m",
	"red":         "\033[31m",
	"green":       "\033[32m",
	"yellow":      "\033[33m",
	"blue":        "\033[34m",
	"magenta":     "\033[35m",
	"cyan":        "\033[36m",
	"white":       "\033[37m",
	"grey":        "\033[90m",
	"default":     "\033[39m",
	"reset":       "\033[0m",
	"bold":        "\033[1m",
	"underline":   "\033[4m",
	"blink":       "\033[5m",
	"reverse":     "\033[7m",
	"hidden":      "\033[8m",
	"blackbg":     "\033[40m",
	"redbg":       "\033[41m",
	"greenbg":     "\033[42m",
	"yellowbg":    "\033[43m",
	"bluebg":      "\033[44m",
	"magentabg":   "\033[45m",
	"cyanbg":      "\033[46m",
	"whitebg":     "\033[47m",
	"resetbg":     "\033[49m",
	"boldbg":      "\033[1m",
	"underlinebg": "\033[4m",
	"blinkbg":     "\033[5m",
	"reversebg":   "\033[7m",
	"hiddenbg":    "\033[8m",
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
