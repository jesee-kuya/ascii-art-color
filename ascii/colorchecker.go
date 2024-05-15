package ascii

import (
	"strings"
)

func ColorValue(value string, Col []string) string {
	if strings.Contains(value, "-bright") {
		return Col[0]
	} else if strings.Contains(value, "-dark") {
		return Col[1]
	} else if strings.Contains(value, "-highlight") {
		return Col[2]
	} else {
		return Col[3]
	}
}

func ColorChecker(color string, data map[string][]string) string {
	color = strings.ToLower(color)
	if strings.Contains(color, "yellow") {
		value := data["yellow"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "red") {
		value := data["red"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "green") {
		value := data["green"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "blue") {
		value := data["blue"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "gray") {
		value := data["gray"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "magenta") {
		value := data["magenta"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "cyan") {
		value := data["cyan"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "white") {
		value := data["white"]
		return ColorValue(color, value)
	} else if strings.Contains(color, "black") {
		value := data["black"]
		return ColorValue(color, value)
	} else {
		return "\033[0m"
	}
}
