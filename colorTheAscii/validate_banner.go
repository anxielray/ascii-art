package colortheascii

import (
	"os"
	"strings"
)

func Validate_banner() (banner string) {
	if len(os.Args) == 5 && strings.HasPrefix(os.Args[1], "--color=") &&
		(strings.HasPrefix(os.Args[4], "standard") ||
			strings.HasPrefix(os.Args[4], "shadow") ||
			strings.HasPrefix(os.Args[4], "thinkertoy")) {
		banner = os.Args[4]
	}
	return
}
