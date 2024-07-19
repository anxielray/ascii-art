package colortheascii

import (
	"bufio"
	"fmt"
	"os"
)

func SecondaryMap(n int, fileName string) string {
	file, err := os.Open(fmt.Sprintf("%s.txt", fileName))
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
				str += string(ch)
			}
		}
		ascMap[n] = str
	}
	defer file.Close()
	return str
}
