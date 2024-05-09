package ascii

import (
	"errors"
	"strings"
)

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
