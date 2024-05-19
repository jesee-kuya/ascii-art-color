package main

import (
	"flag"
	"fmt"

	"ascii/ascii"
)

func main() {
	var filename string
	var colorflag string
	var lettersTocolor []string
	var paint string

	flag.StringVar(&filename, "filename", "standard", "name for the files")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.Parse()
	words := flag.Args()

	bannerContent, err := ascii.GetFileName(filename)
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
