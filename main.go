package main

import (
	"flag"
	"fmt"
	"strings"

	"ascii/ascii"
)

func main() {
	Colormap := map[string]string{
		"yellow":  "\033[33;1m",
		"red":     "\033[31;1m",
		"green":   "\033[32;1m",
		"blue":    "\033[34;1m",
		"magenta": "\033[35;1m",
		"cyan":    "\033[36;1m",
		"gray":    "\033[37;1m",
		"white":   "\033[97;1m",
		"black":   "\033[30;1m",
		"purple":  "\033[38;2;128;0;128m",
		"orange":  "\033[38;2;255;165;0m",
		"brown":   "\033[38;2;165;42;42m",
		"pink":    "\033[38;2;255;192;203m",
		"gold":    "\033[38;2;255;215;0m",
		"silver":  "\033[38;2;192;192;192m",
		"violet":  "\033[38;2;238;130;238m",
		"maroon":  "\033[38;2;128;0;0m",
		"navy":    "\033[38;2;0;0;128m",
		"olive":   "\033[38;2;128;128;0m",
	}

	var filename string
	var colorflag string
	var option []string
	var str string
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
		colorflag = strings.ToLower(colorflag)

		_, ok := Colormap[colorflag]
		if !ok {
			if strings.Contains(colorflag, "rgb") {
				str, err = ascii.Rgb(colorflag)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Printf("The color %v is not yet defined. Try another color.\n", colorflag)
				return
			}
		}
		var paint string
		if str == "" {
			paint = ascii.ColorChecker(colorflag, Colormap)
		} else {
			paint = str
		}

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
