package main

import (
	"flag"
	"fmt"
	"strings"

	"ascii/ascii"
)

func main() {
	yellow := "\033[33;1m"
	red := "\033[31;1m"
	green := "\033[32;1m"
	blue := "\033[34;1m"
	magenta := "\033[35;1m"
	cyan := "\033[36;1m"
	gray := "\033[37;1m"
	white := "\033[97;1m"
	black := "\033[30;1m"

	Colormap := map[string]string{
		"yellow":  yellow,
		"red":     red,
		"green":   green,
		"blue":    blue,
		"magenta": magenta,
		"cyan":    cyan,
		"gray":    gray,
		"white":   white,
		"black":   black,
	}

	var filename string
	var colorflag string
	var option []string
	flag.StringVar(&filename, "filename", "standard", "name for the files")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.Parse()
	words := flag.Args()

	content, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ascii.IsFlagPassed("color") {
		var test string
		colorflag = strings.ToLower(colorflag)

		_, ok := Colormap[colorflag]
		if !ok {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		colorflag = test + colorflag

		paint := ascii.ColorChecker(colorflag, Colormap)

		if len(words) == 1 {
			ascii.Art(words, content, option, paint, 0)
		} else if len(words) == 2 {
			option = append(option, words[0])
			ascii.Art(words, content, option, paint, 1)
		} else {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	} else {
		paint := ""
		ascii.Art(words, content, option, paint, 0)
	}
}
