package ascii

import (
	"errors"
	"strconv"
	"strings"
)

func HexToRgb(color string) (r, g, b uint8, err error) {
	if strings.Contains(color, "=") {
		return 0, 0, 0, errors.New("usage: go run . --color=<color> <letters to be colored> \"something\"")
	}
	// Remove the '#' symbol if present
	color = strings.TrimPrefix(color, "#")
	if len(color) != 6 {
		return 0, 0, 0, errors.New("the hex code entered is not correct")
	}
	// Convert the hex code to decimal
	rStr := color[:2]
	gStr := color[2:4]
	bStr := color[4:]

	rInt, err := strconv.ParseUint(rStr, 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}
	gInt, err := strconv.ParseUint(gStr, 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}
	bInt, err := strconv.ParseUint(bStr, 16, 8)
	if err != nil {
		return 0, 0, 0, err
	}

	return uint8(rInt), uint8(gInt), uint8(bInt), nil
}
