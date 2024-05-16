package ascii

func Art(words []string, bannerContent []string, letterToColor string, paint string, index int) {
	word := Arrange(words[index:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(bannerContent, wordsArr, letterToColor, paint)
}
