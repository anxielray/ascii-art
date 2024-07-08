package colortheascii

import (
	"strings"
)

func IndexTracker(letters, text string) (Index []int) {
	var wanted int
	if strings.Contains(text, letters) {
		wanted = strings.Index(text, letters)
		Index = append(Index, wanted)
		Index = append(Index, IndexTracker(letters, text[(wanted+1):])...)
	}

	return
}

// text = strings.ReplaceAll(text, letters, fmt.Sprintf("\033[1;%dm%s\033[0m", GetColor(color), letters))
