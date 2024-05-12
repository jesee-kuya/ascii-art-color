package ascii_test

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"ascii/ascii"
)

// Ascii testcases
var testCases = []struct {
	name                                                   string
	line1, line2, line3, line4, line5, line6, line7, line8 string
	input                                                  string
}{
	{
		"Test1",
		" _    _          _   _          \n",
		"| |  | |        | | | |         \n",
		"| |__| |   ___  | | | |   ___   \n",
		"|  __  |  / _ \\ | | | |  / _ \\  \n",
		"| |  | | |  __/ | | | | | (_) | \n",
		"|_|  |_|  \\___| |_| |_|  \\___/  \n",
		"                                \n",
		"                                \n",
		"Hello",
	},
	{
		"Test2",
		" _    _          _   _          \n",
		"| |  | |        | | | |         \n",
		"| |__| |   ___  | | | |   ___   \n",
		"|  __  |  / _ \\ | | | |  / _ \\  \n",
		"| |  | | |  __/ | | | | | (_) | \n",
		"|_|  |_|  \\___| |_| |_|  \\___/  \n",
		"                                \n",
		"                                \n\n",
		"Hello\n",
	},
}

// CheckAscii testcases
var testCases1 = []struct {
	name     string
	word     string
	expected bool
}{
	{"Test1", "Hello", true},
	{"Test2", "I like the ¥", false},
	{"Test3", "What a sign §", false},
	{"Test4", "Copyright ©", false},
	{"Test5", "Wrong Turn ⊗", false},
	{"Test6", "Inverted thinkers ⊥", false},
	{"Test7", "I like a game of ♥♣♠", false},
	{"Test8", "Percentile ‰", false},
}

// Arrange testcases
var testCases2 = []struct {
	name     string
	input    []string
	expected string
}{
	{"Test1", []string{"Hello", "there", "you."}, "Hello there you."},
	{"Test2", []string{"How", "are", "you?"}, "How are you?"},
}

// Slice testcases
var testCases3 = []struct {
	name     string
	word     string
	expected []string
}{
	{"Test1", "Hello\nThere", []string{"Hello", "There"}},
}

// NoError testcases
var testCases4 = []struct {
	name     string
	word     []string
	expected bool
}{
	{"Test1", []string{""}, false},
	{"Test2", []string{"Hello", "There"}, true},
}


// TestAscii tests the Ascii function
func TestAscii(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// redirect stdout to pipe
			old := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// read ascii art from a file
			content, _ := ascii.Reader("/home/jkuya/ascii-art/standard.txt", "\n")
			// split input into lines
			input := strings.Split(tc.input, "\n")

			// run the ascii function
			ascii.Ascii(content, input)

			// close the write end of pipe
			w.Close()
			// restore stdout
			os.Stdout = old

			// read from the read end of pipe
			var buf bytes.Buffer
			_, err := buf.ReadFrom(r)
			if err != nil {
				t.Fatalf("Error reading from pipe: %v", err)
			}
			// compare output to expected output
			expected := tc.line1 + tc.line2 + tc.line3 + tc.line4 + tc.line5 + tc.line6 + tc.line7 + tc.line8
			if buf.String() != expected {
				t.Errorf("got \n %v expected \n %v", buf.String(), expected)
			}
		})
	}
}

func TestCheckAscii(t *testing.T) {
	for _, tc := range testCases1 {
		t.Run(tc.name, func(t *testing.T) {
			wordArr := strings.Split(tc.word, "")
			got := ascii.CheckAscii(wordArr)
			if got != tc.expected {
				t.Errorf("got %v \n expected %v \n", got, tc.expected)
			}
		})
	}
}

func TestArrange(t *testing.T) {
	for _, tc := range testCases2 {
		t.Run(tc.name, func(t *testing.T) {
			got := ascii.Arrange(tc.input)
			if got != tc.expected {
				t.Errorf("got %v \nexpected %v \n", got, tc.expected)
			}
		})
	}
}

func TestSlice(t *testing.T) {
	for _, tc := range testCases3 {
		got := ascii.Slice(tc.word)
		word1 := strings.Join(got, " ")
		word2 := strings.Join(tc.expected, " ")
		if word1 != word2 {
			t.Errorf("got %v \n expected %v \n", got, tc.expected)
		}
	}
}

func TestNoError(t *testing.T) {
	for _, tc := range testCases4 {
		t.Run(tc.name, func(t *testing.T) {
			got := ascii.NoError(tc.word)
			if got != tc.expected {
				t.Errorf("got %v \nexpected %v \n", got, tc.expected)
			}
		})
	}
}
