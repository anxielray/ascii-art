package colortheascii

import (
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
