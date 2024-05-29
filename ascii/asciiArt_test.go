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
	lettersToColor                                         string
}{
	{
		"Test1",
		"\033[31;1m _    _  \033[0m\033[31;1m       \033[0m\033[31;1m _  \033[0m\033[31;1m _  \033[0m\033[31;1m        \033[0m\n",
		"\033[31;1m| |  | | \033[0m\033[31;1m       \033[0m\033[31;1m| | \033[0m\033[31;1m| | \033[0m\033[31;1m        \033[0m\n",
		"\033[31;1m| |__| | \033[0m\033[31;1m  ___  \033[0m\033[31;1m| | \033[0m\033[31;1m| | \033[0m\033[31;1m  ___   \033[0m\n",
		"\033[31;1m|  __  | \033[0m\033[31;1m / _ \\ \033[0m\033[31;1m| | \033[0m\033[31;1m| | \033[0m\033[31;1m / _ \\  \033[0m\n",
		"\033[31;1m| |  | | \033[0m\033[31;1m|  __/ \033[0m\033[31;1m| | \033[0m\033[31;1m| | \033[0m\033[31;1m| (_) | \033[0m\n",
		"\033[31;1m|_|  |_| \033[0m\033[31;1m \\___| \033[0m\033[31;1m|_| \033[0m\033[31;1m|_| \033[0m\033[31;1m \\___/  \033[0m\n",
		"\033[31;1m         \033[0m\033[31;1m       \033[0m\033[31;1m    \033[0m\033[31;1m    \033[0m\033[31;1m        \033[0m\n",
		"\033[31;1m         \033[0m\033[31;1m       \033[0m\033[31;1m    \033[0m\033[31;1m    \033[0m\033[31;1m        \033[0m\n",
		"Hello",
		"Hello",
	},
	{
		"Test2",
		"\033[31;1m _    _  \033[0m        _   _  \033[31;1m        \033[0m\n",
		"\033[31;1m| |  | | \033[0m       | | | | \033[31;1m        \033[0m\n",
		"\033[31;1m| |__| | \033[0m  ___  | | | | \033[31;1m  ___   \033[0m\n",
		"\033[31;1m|  __  | \033[0m / _ \\ | | | | \033[31;1m / _ \\  \033[0m\n",
		"\033[31;1m| |  | | \033[0m|  __/ | | | | \033[31;1m| (_) | \033[0m\n",
		"\033[31;1m|_|  |_| \033[0m \\___| |_| |_| \033[31;1m \\___/  \033[0m\n",
		"\033[31;1m         \033[0m               \033[31;1m        \033[0m\n",
		"\033[31;1m         \033[0m               \033[31;1m        \033[0m\n\n",
		"Hello\n",
		"Ho",
	},
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
			fileArr, _ := ascii.Reader("/home/jkuya/ascii-art/standard.txt", "\n")
			// split input into lines
			wordsArr := strings.Split(tc.input, "\n")

			// run the ascii function
			ascii.Ascii(fileArr, wordsArr, tc.lettersToColor, "\033[31;1m")

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
