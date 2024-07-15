package colortheascii

import (
	"fmt"
	"os"
	"strings"
)

func Ascii_art() {
	text := os.Args[1]
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(Art_Concatenator(a, b, c, d, e, f, g, h))
		} else {
			fmt.Println()
		}
	}
}
