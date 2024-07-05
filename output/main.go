package main

import (
	"fmt"
	"os"
	"strings"

	output "output/output"
)

func main() {
	arguments := len(os.Args)
	if arguments < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEx: go run . --output=<fileName.txt> something standard")
		return
	} else if arguments == 2 {
		if os.Args[1] == "" {
			return
		}
		if os.Args[1] == "\\n" {
			fmt.Println()
			return
		}
		input := output.Errors(os.Args[1])
		if strings.Contains(input, "\\t") {
			input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
		}
		lines := strings.Split(input, "\\n")
		for _, line := range lines {
			if line != "" {
				var a, b, c, d, e, f, g, h []int = output.Calculator(line)
				str1, str2, str3, str4, str5, str6, str7, str8 := output.Concatenator(a, b, c, d, e, f, g, h)
				fmt.Print(str1, str2, str3, str4, str5, str6, str7, str8)
			} else if line == "\\n" {
				fmt.Println()
			} else {
				fmt.Println()
			}
		}
		os.Exit(0)
	} else if arguments == 3 && !strings.HasPrefix(os.Args[1], "--output=") {
		if os.Args[1] == "" {
			return
		}
		if os.Args[1] == "\\n" {
			fmt.Println()
			return
		}
		input := output.Errors(os.Args[1])
		if strings.Contains(input, "\\t") {
			input = strings.ReplaceAll(os.Args[1], "\\t", "    ")
		}
		lines := strings.Split(input, "\\n")
		for _, line := range lines {
			if line != "" {
				var a, b, c, d, e, f, g, h []int = output.Calculator(line)
				str1, str2, str3, str4, str5, str6, str7, str8 := output.Concatenatorfs(a, b, c, d, e, f, g, h)
				fmt.Print(str1, str2, str3, str4, str5, str6, str7, str8)
			} else if line == "\\n" {
				fmt.Println()
			} else {
				fmt.Println()
			}
		}
		os.Exit(0)
	} else if arguments == 4 && strings.HasPrefix(os.Args[1], "--output=") {
		output.Ascii_Output(os.Args[1], os.Args[2], os.Args[3])
	} else if len(os.Args) == 3 && strings.HasPrefix(os.Args[1], "--output=") {
		output.LessArguments(os.Args[1], os.Args[2])
	}
}
