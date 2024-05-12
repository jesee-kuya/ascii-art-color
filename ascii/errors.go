package ascii

import (
	"fmt"
	"os"
)

// NoError checks for errors in the input arguments
func NoError(words []string) bool {
	// check if there is no input passed and flag without input
	if len(words) <= 1 {
		fmt.Println("Error: usage; go run . <input>")
		return false
	}
	// check for incorrect flag usage
	if (os.Args[1] == "-t" || os.Args[1] == "-s") && len(words) <= 2 {
		fmt.Println("Error: usage; go run . <flag> <input>")
		return false
	}
	// check for empty inputs
	if os.Args[1] == "" && len(os.Args) == 2 {
		return false
	} else if os.Args[1] == "-t" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	} else if os.Args[1] == "-s" && os.Args[2] == "" && len(os.Args) == 3 {
		return false
	}
	return true
}
