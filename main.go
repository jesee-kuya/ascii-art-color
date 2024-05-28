package main

import (
	"flag"
	"fmt"
	"os"
	"ascii/ascii"
)
// main function is used to read the banner file and print the ascii art
// based on the arguments passed
func main() {
	for _ , v := range os.Args {
		if v == "--color" || v == "-color" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
	}
	var filename string
	var colorflag string
	var lettersTocolor string
	var colorCode string

	flag.StringVar(&filename, "filename", "standard", "name for the files")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.Parse()
	 argsPassed:= flag.Args()

	bannerContent, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ascii.IsFlagPassed("color") {
		ascii.Color(colorflag, lettersTocolor, argsPassed, bannerContent)
	} else {
		colorCode = ""
		ascii.Art(argsPassed, bannerContent, lettersTocolor, colorCode, 0)
	}
}
