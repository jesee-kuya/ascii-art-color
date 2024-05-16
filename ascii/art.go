package ascii

func Art(words []string, content []string, option []string, paint string, index int) {
	word := Arrange(words[index:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(content, wordsArr, option, paint)
}
