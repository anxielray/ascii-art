package colortheascii

import (
	"bufio"
	"fmt"
	"os"
)

func SecondaryMap(n int, color string) string {
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
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
				str += ColorText(color, string(ch))
			}
		}
		ascMap[n] = str
	}
	defer file.Close()
	return str
}
