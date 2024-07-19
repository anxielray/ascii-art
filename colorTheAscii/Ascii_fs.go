package colortheascii

import (
	"fmt"
	"os"
	"strings"
)

func Ascii_fs() {
	var banner string
	if strings.HasSuffix(os.Args[2], ".txt") {
		banner = strings.TrimSuffix(os.Args[2], ".txt")
	} else {
		banner = os.Args[2]
	}
	text := os.Args[1]
	if strings.Contains(text, "\\t") {
		text = strings.ReplaceAll(text, "\\t", "    ")
	}
	lines := strings.Split(text, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Collector(line)
			fmt.Println(Fs_Concatenator(a, b, c, d, e, f, g, h, banner))
		} else {
			fmt.Println()
		}
	}
}
