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
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "\\n" {
		fmt.Println()
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "" {
		os.Exit(0)
	}
	flag.Parse()
	co.ValidFlag()
	color := *colorFlag
	arguments = flag.Args()
	if len(os.Args) == 2 && flag.NFlag() < 1 {
		co.Ascii_art()
	} else if len(os.Args) == 3 && flag.NFlag() < 1 {
		co.Ascii_fs()
	} else if len(arguments) == 2 {
		co.Ascii_Color()
	} else if len(arguments) == 1 && flag.NFlag() == 1 {
		text = strings.Join(arguments, "")
		co.Ascii_Color_less_Arguments(text, color)
	} else if len(arguments) == 3 {
		co.Full_Ascii_color(color)
	}
}
