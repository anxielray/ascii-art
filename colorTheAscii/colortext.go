package colortheascii

import "fmt"

func ColorText(color, text string) (coloredText string) {
	if color == "orange" || color == "pink" || color == "purple" ||
		color == "#ffa500" || color == "#800080" || color == "#ffc0cb" {
		coloredText = fmt.Sprintf("\033[38;5;%dm%s\033[0m", GetColor(color), text)
	} else {
		coloredText = fmt.Sprintf("\033[1;%dm%s\033[0m", GetColor(color), text)
	}
	return
}
