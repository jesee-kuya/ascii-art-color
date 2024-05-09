package ascii

import (
	"errors"
	"os"
	"strings"
)

// Reader reads the banner file and returns a slice of string
func Reader(filename string, sepp string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error reading banner file")
	}

// check if the bannerfile has all the values
	content := strings.Split(string(file), sepp)
	if len(content) != 856 {
		return nil, errors.New("the banner file content is not correct")
	}
	return content, nil
}

