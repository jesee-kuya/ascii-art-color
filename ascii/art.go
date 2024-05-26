package ascii


// Art generates ASCII art from a given slice of words and prints it to the console.
// It takes in four arguments: argsPassed (a slice of strings), bannerFileContent (a slice of strings),
// lettersToColor (a string), and paint (a string). The function returns nothing.
func Art(argsPassed []string, bannerFileContent []string, lettersToColor string, colorCode string, indexToStartDisplay int) {
	word := Arrange(argsPassed[indexToStartDisplay:])
	wordsToDisplay := Slice(word)
	if !CheckAscii(wordsToDisplay) {
		return
	}
	Ascii(bannerFileContent, wordsToDisplay, lettersToColor, colorCode)
}
