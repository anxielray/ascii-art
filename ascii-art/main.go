package main

import (
	"fmt"
	"os"
	"strings"

	art "ascii/ascii"
)

func main() {
	arguments := len(os.Args)
	if arguments != 2 {
		fmt.Println("Wrong passing of arguments")
		return
	}
	if os.Args[1] == "" {
		return
	}
	if os.Args[1] == "\\n" {
		fmt.Println()
		return
	}
	input := art.Errors(os.Args[1])
	if strings.Contains(input, "\\t") {
		input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
	}
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = art.Calculator(line)
			str1, str2, str3, str4, str5, str6, str7, str8 := art.Concatenator(a, b, c, d, e, f, g, h)
			fmt.Println(str1, str2, str3, str4, str5, str6, str7, str8)
		} else if line == "\\n" {
			fmt.Println()
		} else {
			fmt.Println()
		}
	}
	os.Exit(0)
}
