package ascii

// Arrange converts a slice of strings into a single string.
func Arrange(words []string) string {
	var word string
	for i, v := range words {
		// If it's not the first word, add a space before the word.
		if i != 0 {
			word += " " + string(v)
		} else {
			word += string(v)
		}
	}
	return word
}
