package main

import (
	"fmt"
	"strings"
)

func LettersColor(text, letters, color string) {
	var index int
	if strings.Contains(text, letters) {
		index = strings.Index(text, letters)
		// text = strings.ReplaceAll(text, letters, fmt.Sprintf("\033[1;%dm%s\033[0m", GetColor(color), letters))
		// var a, b, c, d, e, f, g, h []int = Collector(text)
	}

	fmt.Println(Concatenator(IndexTracker(letters, text), a, b, c, d, e, f, g, h, color, letters))
}
