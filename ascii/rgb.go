package ascii

import (
	"errors"
	"strconv"
	"strings"
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

func Rgb(colorflag string) (string, error) {
	rgb := RGB{}
	colorflag = strings.Trim(colorflag, "rgb")
	colorflag = strings.Trim(colorflag, "()")
	arr := strings.Split(colorflag, ",")
	if len(arr) != 3 {
		return "", errors.New("the rgb code entered is not correctly formatted")
	}
	for i, v := range arr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return "", errors.New("the value entered as rgb code is not correct")
		}
		if i == 0 && num <= 255 {
			rgb.R = num
		} else if i == 1 && num <= 255 {
			rgb.G = num
		} else if i == 2 && num <= 255 {
			rgb.B = num
		} else {
			return "", errors.New("the rgb code is wrong")
		}
	}
	str := RgbToAnsi(rgb)
	return str, nil
}
