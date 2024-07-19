package colortheascii

import (
	"fmt"
	"os"
	"strings"
)

// the error function will validate the command line arguments, their parsing has to be in the right order.
func Error() {
	if len(os.Args) == 1 || len(os.Args) > 5 {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "\\n" {
		fmt.Println()
		os.Exit(0)
	} else if len(os.Args) == 2 && (os.Args[1] == "" || strings.HasSuffix(os.Args[1], ".txt")) {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
		os.Exit(0)
	} else if len(os.Args) == 2 && strings.HasPrefix(os.Args[1], "--color=") {
		fmt.Println(`Usage: go run . [OPTION] [STRING]
			
EX: go run . --color=<color> <substring to be colored> "something"`)
		os.Exit(0)
	}
}
