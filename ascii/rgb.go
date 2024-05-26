package ascii

import (
	"strconv"
)

type RGB struct {
	R int
	G int
	B int
}

func RgbToAnsi(rgb RGB) string {
	str := "\033[38;2;" + strconv.FormatInt(int64(rgb.R), 10) + ";" + strconv.FormatInt(int64(rgb.G), 10) + ";" + strconv.FormatInt(int64(rgb.B), 10) + "m"
	return str
}
