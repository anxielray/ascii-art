package colortheascii

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func ValidFlag() {
	if strings.HasPrefix(os.Args[1], "--color") {
		if !strings.Contains(os.Args[1], "=") {
			fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
			os.Exit(0)
		}
	}
}

func Checker() {
	var (
		colorFlag = flag.String("color", "", "specify the color... e.g, (--color=red)")
		arguments []string
		text      string
	)
	flag.Parse()
	color := *colorFlag
	ValidFlag()
	arguments = flag.Args()

	// ascii-art
	if len(os.Args) == 2 && !strings.HasPrefix(os.Args[1], "--color=") {
		Ascii_art()

		// ascii-fs (with enough or less arguments)
	} else if len(os.Args) == 3 && !strings.Contains(os.Args[1], "--color=") && !strings.HasSuffix(os.Args[1], ".txt") {
		Ascii_fs()
	} else if len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--color=") && (!strings.HasPrefix(os.Args[2], "standard") && !strings.HasPrefix(os.Args[1], "shadow") && !strings.HasPrefix(os.Args[1], "thinkertoy")) {
		text = os.Args[2]
		Ascii_Color_less_Arguments(text, color)
	} else if (len(arguments) == 2 && flag.NFlag() == 1) && (strings.HasPrefix(os.Args[3], "shadow") ||
		strings.HasPrefix(os.Args[3], "standard") || strings.HasPrefix(os.Args[3], "thinkertoy")) {
		Less_full(color)
	} else if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--color=") && (!strings.HasPrefix(os.Args[1], "standard") && !strings.HasPrefix(os.Args[1], "shadow") && !strings.HasPrefix(os.Args[1], "thinkertoy")) {
		Ascii_Color()
	} else if len(arguments) == 3 {
		Full_Ascii_color(color)
	}
}
