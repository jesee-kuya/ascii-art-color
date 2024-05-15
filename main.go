package main

import (
	"flag"
	"fmt"

	"ascii/ascii"
)

func main() {
	yellow := []string{"\033[33;1m", "\033[33;2m", "\033[33;7m", "\033[33m"}
	red := []string{"\033[31;1m", "\033[31;2m", "\033[31;7m", "\033[31m"}
	green := []string{"\033[32;1m", "\033[32;2m", "\033[32;7m", "\033[32m"}
	blue := []string{"\033[34;1m", "\033[34;2m", "\033[34;7m", "\033[34m"}
	magenta := []string{"\033[35;1m", "\033[35;2m", "\033[35;7m", "\033[35m"}
	cyan := []string{"\033[36;1m", "\033[36;2m", "\033[36;7m", "\033[36m"}
	gray := []string{"\033[37;1m", "\033[37;2m", "\033[37;7m", "\033[37m"}
	white := []string{"\033[97;1m", "\033[97;2m", "\033[97;7m", "\033[97m"}
	black := []string{"\033[30;1m", "\033[30;2m", "\033[30;7m", "\033[30m"}

	Colormap := map[string][]string{
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
	flag.StringVar(&colorflag, "color", "normal", "color for color input")
	flag.Parse()
	words := flag.Args()

	content, err := ascii.GetFileName(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	if ascii.IsFlagPassed("color") {

		paint := ascii.ColorChecker(colorflag, Colormap)

		if len(words) == 1 {
			word := ascii.Arrange(words)
			wordsArr := ascii.Slice(word)
			if !ascii.CheckAscii(wordsArr) {
				return
			}
			ascii.Ascii(content, wordsArr, option, paint)
		} else if len(words) == 2 {
			option = append(option, words[0])
			word := ascii.Arrange(words[1:])
			wordsArr := ascii.Slice(word)
			if !ascii.CheckAscii(wordsArr) || !ascii.CheckAscii(option) {
				return
			}
			ascii.Ascii(content, wordsArr, option, paint)

		} else {
			fmt.Println("error with number of arguments")
			return
		}
	} else {
		paint := ""
		word := ascii.Arrange(words)
		wordsArr := ascii.Slice(word)
		if !ascii.CheckAscii(wordsArr) {
			return
		}
		ascii.Ascii(content, wordsArr, option, paint)
	}
}
