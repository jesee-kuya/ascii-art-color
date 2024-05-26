package ascii

import (
	"errors"
	"strings"
)

// GetFileName retrieves the content of a specified ASCII art file.
// It takes the name of the file as a parameter and returns the content of the file
// as a slice of strings and an error if the file is not found or cannot be read.
func GetFileName(name string) ([]string, error) {
	if strings.ToLower(name) == "thinkertoy" || strings.ToLower(name) == "thinkertoy.txt" {
		content, error := Reader("thinkertoy.txt", "\r\n")
		if error != nil {
			return nil, error
		}
		return content, error
	} else if strings.ToLower(name) == "shadow" || strings.ToLower(name) == "shadow.txt" {
		content, error := Reader("shadow.txt", "\n")
		if error != nil {
			return nil, error
		}
		return content, error
	} else if strings.ToLower(name) == "standard" || strings.ToLower(name) == "standard.txt" {
		content, error := Reader("standard.txt", "\n")
		if error != nil {
			return nil, error
		}
		return content, error
	} else {
		return nil, errors.New("the file is not available or missing")
	}
}
