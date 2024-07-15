package colortheascii

import (
	"fmt"
	"os"
	"strings"
)

func Full_Ascii_color(color string) {
	text := os.Args[3]
	letters := os.Args[2]
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
