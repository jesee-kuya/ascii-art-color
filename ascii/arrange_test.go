package ascii_test

import (
	"testing"

	"ascii/ascii"
)

var testCases = []struct {
    testName string
    data     []string
    expected string
}{
    {
        testName: "Empty Input Slice",
        data:     []string{},
        expected: "",
    },
    {
        testName: "Slice With One Word",
        data:     []string{"hello"},
        expected: "hello",
    },
    {
        testName: "Slice With Multiple Words",
        data:     []string{"hello", "world"},
        expected: "hello world",
    },
}

func TestArrange(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			got := ascii.Arrange(tc.data)
			if got != tc.expected {
				t.Errorf("got %v \nexpected %v \n", got, tc.expected)
			}
		})
	}
}
