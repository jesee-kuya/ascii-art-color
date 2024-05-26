package ascii

import (
	"testing"
)

var testCases = []struct {
	testName              string
    argsPassed            []string
    bannerFileContent     []string
    lettersToColor        string
    colorCode             string
    indexToStartDisplay   int
}{
	{
		testName: "Empty Input Slice",
		argsPassed: []string{},
		bannerFileContent: []string{},
		lettersToColor: "",
		colorCode: "",
		indexToStartDisplay: 0,
	},
	{
		testName: "Slice With One Word",
		argsPassed: []string{"hello"},
		bannerFileContent: []string{"hello"},
		lettersToColor: "ll",
		colorCode: "red",
		indexToStartDisplay: 0,
	},
	{
		testName: "Slice With Multiple Words",
		argsPassed: []string{"hello", "world"},
		bannerFileContent: []string{"hello", "world"},
		lettersToColor: "ll",
		colorCode: "red",
		indexToStartDisplay: 0,
	},
	{
        testName:            "Valid Input",
        argsPassed:          []string{"hello", "world"},
        bannerFileContent:   []string{"  o  ", " /\\ ", "/  \\"},
        lettersToColor:      "o",
        colorCode:           "red",
        indexToStartDisplay: 0,
    },
}

	func TestArt(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.testName, func(t*testing.T) {)
		}
	}
