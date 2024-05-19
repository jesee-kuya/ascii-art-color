package ascii

// Art takes words a slice of string from a certain index, joins the words,
// checks if there is a non-ascii character
// and calls the Ascii function to print the words in ascii-art
func Art(words []string, content []string, option []string, paint string, index int) {
	word := Arrange(words[index:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(bannerContent, wordsArr, letterToColor, paint)
}
