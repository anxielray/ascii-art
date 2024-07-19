package colortheascii

import (
	"fmt"
	"os"
	"strings"
)

func Full_Ascii_color(color string) {
	banner := os.Args[4]
	if len(os.Args) == 5 && strings.HasPrefix(os.Args[1], "--color=") {
		if strings.HasSuffix(os.Args[4], ".txt") {
			banner = strings.TrimSuffix(os.Args[4], ".txt")
		} else {
			banner = os.Args[4]
		}
	}
	text := os.Args[3]
	letters := os.Args[2]
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(Full_concatenator(IndexTracker(letters, text), a, b, c, d, e, f, g, h, color, letters, banner))
		} else {
			fmt.Println()
		}
	}
}

func Less_full(color string) {
	banner := os.Args[3]
	if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--color=") {
		if strings.HasSuffix(os.Args[3], ".txt") {
			banner = strings.TrimSuffix(os.Args[3], ".txt")
		} else {
			banner = os.Args[3]
		}
	}
	text := os.Args[2]
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(Less_Concat_Full(a, b, c, d, e, f, g, h, color, banner))
		} else {
			fmt.Println()
		}
	}
}
