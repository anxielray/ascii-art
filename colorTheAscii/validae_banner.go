package colortheascii

import (
	"flag"
	"os"
	"strings"
)

func Validate_banner() (banner string) {
	if len(os.Args) == 4 && flag.NFlag() == 1 &&
		(strings.HasPrefix(os.Args[3], "standard") ||
			strings.Contains(os.Args[3], "shadow") ||
			strings.Contains(os.Args[3], "thinkertoy")) {
		banner = os.Args[3]
	} else {
		Ascii_Color()
	}
	if (len(os.Args) == 3 && flag.NFlag() < 1) || len(os.Args) == 5 {
		if len(os.Args) == 3 && flag.NFlag() < 1 {
			if strings.HasSuffix(os.Args[2], ".txt") {
				banner = strings.TrimSuffix(os.Args[2], ".txt")
			} else {
				banner = os.Args[2]
			}
		} else if len(os.Args) == 5 && flag.NFlag() == 1 {
			if strings.HasSuffix(os.Args[4], ".txt") {
				banner = strings.TrimSuffix(os.Args[4], ".txt")
			} else {
				banner = os.Args[4]
			}
		}
	} else {
		banner = "standard"
	}
	return
}
