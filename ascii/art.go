package ascii

func Art(check []string, content []string, flag string, letter string, n int) {
	word := Arrange(check[n:])
	wordsArr := Slice(word)
	if !CheckAscii(wordsArr) {
		return
	}
	Ascii(content, wordsArr, flag, letter)
}
