package output

import (
	"fmt"
	"os"
	"strings"
)

func Ascii_Output(fileFlag, argument, bannerfile string) {
	var newFile, trim string
	trim = fileFlag[9:]
	if strings.HasSuffix(trim, ".txt") {
		newFile = trim[:len(trim)-4]
	} else {
		newFile = trim
	}
	theFile := fmt.Sprintf("%s.txt", newFile)
	file, err := os.Create(theFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	input := Errors(argument)
	if strings.Contains(input, "\\t") {
		input = strings.ReplaceAll(argument, "\\t", "    ")
	}
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Calculator(line)
			str1, str2, str3, str4, str5, str6, str7, str8 := Concatenatorout(a, b, c, d, e, f, g, h, bannerfile)
			WriteToFile(file, str1, str2, str3, str4, str5, str6, str7, str8)
		} else if line == "\\n" {
			newline(file)
		} else {
			newline(file)
		}
	}
	os.Exit(0)
}

func WriteToFile(createdFile *os.File, line1, line2, line3, line4, line5, line6, line7, line8 string) {
	_, err := createdFile.WriteString(line1)
	_, err = createdFile.WriteString(line2)
	_, err = createdFile.WriteString(line3)
	_, err = createdFile.WriteString(line4)
	_, err = createdFile.WriteString(line5)
	_, err = createdFile.WriteString(line6)
	_, err = createdFile.WriteString(line7)
	_, err = createdFile.WriteString(line8)
	_, err = createdFile.WriteString("\n")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func LessArguments(fileFlag, argument string) {
	file, err := os.Create(strings.Trim(fileFlag, "--output="))
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	input := Errors(argument)
	if strings.Contains(input, "\\t") {
		input = strings.ReplaceAll(argument, "\\t", "    ")
	}
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			var a, b, c, d, e, f, g, h []int = Calculator(line)
			str1, str2, str3, str4, str5, str6, str7, str8 := Concatenator(a, b, c, d, e, f, g, h)
			WriteToFile(file, str1, str2, str3, str4, str5, str6, str7, str8)
		} else if line == "\\n" {
			fmt.Println()
		} else {
			fmt.Println()
		}
	}
	os.Exit(0)
}

func newline(file *os.File) {
	_, err := file.WriteString("\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
