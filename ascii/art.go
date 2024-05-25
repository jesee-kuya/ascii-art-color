package ascii


// Art generates ASCII art from a given slice of words and prints it to the console.
// It takes in four arguments: words (a slice of strings), content (a slice of strings),
// option (a slice of strings), and paint (a string). The function returns nothing.
func Art(words []string, bannerFileContent []string, lettersToColor string, paint string, index int) {
	word := Arrange(words[index:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(bannerFileContent, wordsArr, lettersToColor, paint)
}
