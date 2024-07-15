package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	co "color/colorTheAscii"
)

var (
	colorFlag = flag.String("color", "", "specify the color... e.g, (--color=red)")
	arguments []string
	letters   string
	text      string
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 4 {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
		os.Exit(0)
	}else if len(os.Args) == 2 && os.Args[1] == "\\n" {
		fmt.Println()
		os.Exit(0)
	}
	flag.Parse()
	co.ValidFlag()
	color := *colorFlag
	arguments = flag.Args()
	if len(arguments) == 2 {
		letters = arguments[0]
		text = strings.Join(arguments[1:], "")
		if strings.Contains(text, "\\t") {
			text = strings.ReplaceAll(text, "\\t", "    ")
		}
		var a, b, c, d, e, f, g, h []int = co.Collector(text)
		fmt.Println(co.Concatenator(co.IndexTracker(letters, text), a, b, c, d, e, f, g, h, color, letters))
	} else if len(arguments) == 1 {
		text = strings.Join(arguments, "")
		if strings.Contains(text, "\\t") {
			text = strings.ReplaceAll(text, "\\t", "    ")
		}
		lines := strings.Split(text, "\\n")
		for _, line := range lines {
			if line != "" {
				var a, b, c, d, e, f, g, h []int = co.Collector(text)
				fmt.Println(co.PlainConcatenator(a, b, c, d, e, f, g, h, color))
			} else {
				fmt.Println()
			}
		}
	}
}
