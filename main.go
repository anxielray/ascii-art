package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	colorFlag = flag.String("color", "", "Specify the color (e.g., --color=red)")
	arguments []string
	letters   string
	text      string
)

func main() {
	flag.Parse()

	if *colorFlag == "" || len(flag.Args()) < 1 {
		printUsage()
		return
	}
	color := *colorFlag
	arguments = flag.Args()
	if len(arguments) == 2 {
		letters = arguments[0]
		text = strings.Join(arguments[1:], "")
		colorLetters(color, letters, text)
	} else if len(arguments) == 1 {
		text = strings.Join(arguments, "")
		colorText(color, text)
	}
}

func colorText(color, text string) {
	fmt.Printf("\033[1;%dm%s\033[0m\n", getColor(color), text)
}

func colorLetters(color, letters, text string) {
	var coloredText strings.Builder
	for _, ch := range text {
		if strings.ContainsRune(letters, ch) {
			coloredText.WriteString(fmt.Sprintf("\033[1;%dm%c\033[0m", getColor(color), ch))
		} else {
			coloredText.WriteRune(ch)
		}
	}
	fmt.Println(coloredText.String())
}

func getColor(color string) int {
	switch color {
	case "red":
		return 31
	case "green":
		return 32
	case "yellow":
		return 33
	case "blue":
		return 34
	case "magenta":
		return 35
	case "cyan":
		return 36
	default:
		return 0
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: go run . [OPTION] [STRING]\n\n")
	fmt.Fprintf(os.Stderr, "example: go run . --color=red hello\n")
	flag.PrintDefaults()
}
