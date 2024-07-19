package colortheascii

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func Ascii_Color() {
	text := flag.Args()[1]
	letters := flag.Args()[0]
	color := strings.TrimPrefix(os.Args[1], "--color=")
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(Concatenator(IndexTracker(letters, text), a, b, c, d, e, f, g, h, color, letters))
		} else {
			fmt.Println()
		}
	}
}

//Deals with the scenario when less number of arguments are passed. Without the substring
func Ascii_Color_less_Arguments(text, color string) {
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(PlainConcatenator(a, b, c, d, e, f, g, h, color))
		} else {
			fmt.Println()
		}
	}
}
