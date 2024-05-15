package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ascii "ascii/ascii"
)

var Color = []string{
	"black", "\033[30m",
	"red", "\033[31m",
	"green", "\033[32m",
	"yellow", "\033[33m",
	"blue", "\033[34m",
	"magenta", "\033[35m",
	"cyan", "\033[36m",
	"white", "\033[37m",
	"grey", "\033[90m",
	"default", "\033[39m",
	"reset", "\033[0m",
	"bold", "\033[1m",
	"underline", "\033[4m",
	"blink", "\033[5m",
	"reverse", "\033[7m",
	"hidden", "\033[8m",
	"blackbg", "\033[40m",
	"redbg", "\033[41m",
	"greenbg", "\033[42m",
	"yellowbg", "\033[43m",
	"bluebg", "\033[44m",
	"magentabg", "\033[45m",
	"cyanbg", "\033[46m",
	"whitebg", "\033[47m",
	"resetbg", "\033[49m",
	"boldbg", "\033[1m",
	"underlinebg", "\033[4m",
	"blinkbg", "\033[5m",
	"reversebg", "\033[7m",
	"hiddenbg", "\033[8m",
}

func main() {
	colorFlag := flag.String("color", "reset", "color the command line output")
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
	check := flag.Args()

	if ascii.IsFlagPassed("color") {
		*colorFlag = strings.ToLower(*colorFlag)
		test := false

		for _, v := range Color {
			if strings.Contains(*colorFlag, v) && v != "" {
				*colorFlag = v
				test = true
				break
			}
		}

		if !test && *colorFlag == "" {
			*colorFlag = "reset"
			test = true
		}

		if !test {
			fmt.Printf("The color %v is not yet defined, Try another color.\n", *colorFlag)
			return
		}

		if len(check) < 1 {
			fmt.Println("Usage: go run . --color=<color> <letters to be colored>")
			return
		}

		letter := check[0]

		if len(check) > 1 {
			ascii.Art(check, content, *colorFlag, letter, 1)
		} else {
			ascii.Art(check, content, *colorFlag, letter, 0)
		}
	} else {
		letter := ""
		ascii.Art(check, content, *colorFlag, letter, 0)
	}
}
