package ascii_test

import (
	"testing"

	"ascii/ascii"
)

var testcase = []struct {
	name     string
	data     []string
	expected string
}{
	{"Test1", []string{"Hello", "there", "you."}, "Hello there you."},
	{"Test2", []string{"How", "are", "you?"}, "How are you?"},
}

func TestArrange(t *testing.T) {
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			got := ascii.Arrange(tc.data)
			if got != tc.expected {
				t.Errorf("got %v \nexpected %v \n", got, tc.expected)
			}
		})
	}
}
