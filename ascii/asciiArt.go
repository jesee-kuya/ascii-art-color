package ascii

import (
	"fmt"
	"strings"
)

// Ascii prints ASCII art from a given array of characters.
// The characters are extracted from a predefined file array.
// The function takes in four arguments: fileArr (a slice of strings representing the file array),
// wordsArr (a slice of strings representing the words to be printed),
// lettersToColor (a string representing the letters to be colored),
// and color (a string representing the color to be applied).

func Ascii(bannerFileContent []string, wordsToDisplay []string, lettersToColor string, colorCode string) {
	var count int
	reset := "\033[0m" // color reset code

	for _, val := range wordsToDisplay {
		if val != "" {
			for i := 1; i <= 8; i++ {
				for _, v := range val {
					start := (v - 32) * 9
					// If lettersToColor is empty, apply the color to the character
					if len(lettersToColor) == 0 {
						fmt.Print(colorCode + bannerFileContent[int(start)+i] + reset)

					// If the character is in lettersToColor, apply the color
					} else if strings.Contains(lettersToColor, string(v)) {
						fmt.Print(colorCode + bannerFileContent[int(start)+i] + reset)

						// Otherwise, print the character without the color
					} else {
						fmt.Print(bannerFileContent[int(start)+i])
					}
				}
				fmt.Println()
			}
		} else {
			count++
			if count < len(wordsToDisplay) {
				fmt.Println()
			}
		}
	}
}
