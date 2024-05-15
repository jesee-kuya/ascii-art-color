package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ascii "ascii/ascii"
)

var Color = []string{"black", "red", "green", "yellow", "blue", "magenta", ""}

func main() {
	colorFlag := flag.String("color", "reset", "color the command line output")
	// outPutFlag := flag.String("output", "", "print the output in a file")
	flag.Parse()
	words := os.Args
	if !ascii.NoError(words) {
		return
	}

	content, error := ascii.Reader("standard.txt", "\n")
	if error != nil {
		fmt.Println(error)
		return
	}
	colorize := false
	*colorFlag = strings.ToLower(*colorFlag)
	for _, v := range Color {
		if strings.Contains(*colorFlag, v) && v != "" {
			*colorFlag = v
			colorize = true
		} else if strings.Contains(*colorFlag, v) && v == "" {
			*colorFlag = "reset"
			colorize = true
		}
	}
	switch colorize {
	case true:
		check := flag.Args()
		letter := check[0]
		if len(check) > 1 {
			word := ascii.Arrange(check[1:])
			wordsArr := ascii.Slice(word)
			if !ascii.CheckAscii(wordsArr) {
				return
			}
			ascii.Ascii(content, wordsArr, *colorFlag, letter)
		} else {
			word := ascii.Arrange(check[0:])
			wordsArr := ascii.Slice(word)
			if !ascii.CheckAscii(wordsArr) {
				return
			}
			ascii.Ascii(content, wordsArr, *colorFlag, letter)
		}
	default:
		letter := ""
		word := ascii.Arrange(words[1:])
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr, *colorFlag, letter)
	}
}
