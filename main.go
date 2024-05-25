package main

import (
	"flag"
	"fmt"
	"strings"

	"ascii/ascii"
)

// Colormap defines the mapping of color names to their respective ANSI escape codes.
var Colormap = map[string]string{
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

// Define command-line flags
func main() {

	var filename string
	var colorflag string
	var lettersToColor string
	var str string
	var paint string

	// Define command-line flags
	flag.StringVar(&filename, "filename", "standard", "name for the files")
	flag.StringVar(&colorflag, "color", "reset", "color for color input")
	flag.Parse()
	words := flag.Args()

	// Get the content of the banner file
	bannerFileContent, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if the flag is passed in the command line and process the color input
	if ascii.IsFlagPassed("color") {
		colorflag = strings.ToLower(colorflag)

		// Check if the color is valid
		_, ok := Colormap[colorflag]
		if !ok {
			if strings.Contains(colorflag, "rgb") {
				str, err = ascii.RgbToAnsiConv(colorflag)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Printf("The color %v is not yet defined. Try another color.\n", colorflag)
				return
			}
		}

		// Check if the letters to be colored is passed in the command line
		if str == "" {
			paint = ascii.ColorChecker(colorflag, Colormap)
		} else {
			paint = str
		}

		// Generate ASCII art
		if len(words) == 1 {
			ascii.Art(words, bannerFileContent, lettersToColor, paint, 0)
		} else if len(words) == 2 {
			lettersToColor += words[0]
			ascii.Art(words, bannerFileContent, lettersToColor, paint, 1)
		} else {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored> \"something\"")
			return
		}
		// If the flag is not passed print the default color
	} else {
		paint := ""
		ascii.Art(words, bannerFileContent, lettersToColor, paint, 0)
	}
}
