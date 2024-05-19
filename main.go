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
		ascii.Color(colorflag, lettersTocolor, words, bannerContent)
	} else {
		paint = ""
		ascii.Art(words, bannerContent, lettersTocolor, paint, 0)
	}
}
