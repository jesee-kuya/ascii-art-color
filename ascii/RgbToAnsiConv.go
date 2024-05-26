package ascii

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type RGB struct {
	R int
	G int
	B int
}



// RgbToAnsiConv converts an RGB color to an ANSI escape sequence.
// The ANSI escape sequence is used to set the foreground color of the terminal.
// The function takes a string representing an RGB color in the format "rgb(R, G, B)"
// where R, G, and B are integers between 0 and 255.
// It returns a string that represents the ANSI escape sequence.
func RgbToAnsiConv(colorflag string) (string, error) {
	// Trim the colorflag of any strings that start with "rgb".
	colorflag = strings.Trim(colorflag, "rgb")

	// Trim the colorflag of any parentheses, brackets, and square brackets.
	colorflag = strings.Trim(colorflag, "(){}[]")

	// Trim the colorflag of any whitespace.
	colorflag = strings.ReplaceAll(colorflag, " ", "")

	arr := strings.Split(colorflag, ",")
	if len(arr) != 3 {
		return "", errors.New("use rgb(r, g, b) format")
	}

	// Create an RGB struct.
	rgb := RGB{}
	for i, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return "", errors.New("the rgb code is wrong. use rgb(r, g, b) format")
		}
		// Assign the integer to the corresponding field of the RGB struct.
		if i == 0 && num <= 255 {
			rgb.R = num
		} else if i == 1 && num <= 255 {
			rgb.G = num
		} else if i == 2 && num <= 255 {
			rgb.B = num
		} else {
			return "", errors.New("the rgb code is wrong. use rgb(r, g, b) format")
		}
	}
	// Convert the RGB struct to an ANSI escape sequence and return it.
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", rgb.R, rgb.G, rgb.B), nil
}
