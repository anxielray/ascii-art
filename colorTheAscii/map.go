package colortheascii

import (
	"bufio"
	"fmt"
	"os"
)

func Maps(n int) string {
	banner := Validate_banner()
	file, err := os.Open(fmt.Sprintf("%s.txt", banner))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	scanner := bufio.NewScanner(file)
	var count int
	var line, str string
	for scanner.Scan() {
		line = scanner.Text()
		count++
		ascMap := make(map[int]string)
		if count == n {
			for _, ch := range line {
				str += string(ch)
			}
		}
		ascMap[n] = str
	}
	defer file.Close()
	return str
}
