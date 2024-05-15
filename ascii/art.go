package ascii

func Art(words []string, content []string, option []string, paint string, n int) {
	word := Arrange(words[n:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(content, wordsArr, option, paint)
}
