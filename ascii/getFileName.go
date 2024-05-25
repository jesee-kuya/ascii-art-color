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
		bannerFileContent, error := Reader("thinkertoy.txt", "\r\n")
		if error != nil {
			return nil, error
		}
		return bannerFileContent, error
	} else if strings.ToLower(name) == "shadow" || strings.ToLower(name) == "shadow.txt" {
		bannerFileContent, error := Reader("shadow.txt", "\n")
		if error != nil {
			return nil, error
		}
		return bannerFileContent, error
	} else if strings.ToLower(name) == "standard" || strings.ToLower(name) == "standard.txt" {
		bannerFileContent, error := Reader("standard.txt", "\n")
		if error != nil {
			return nil, error
		}
		return bannerFileContent, error
	} else {
		return nil, errors.New("error: The file is not available or missing")
	}
}
