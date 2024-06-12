package colortheascii

import "fmt"

func ColorText(color, text string) (coloredText string) {
	coloredText = fmt.Sprintf("\033[1;%dm%s\033[0m", GetColor(color), text)
	return
}
