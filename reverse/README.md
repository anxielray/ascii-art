# _*ASCII-ART-REVERSE*_

- A simple Go program to reverse the ASCII graphic representation into text.

## _*Description*_

Ascii-reverse consists on reversing the process, converting the graphic representation into a text. The program takes a flag as an argument, `--reverse=<fileName>`, where `--reverse` is the flag and `<fileName>` is the file name. The program then prints the ASCII art in normal text.

- If any other formats are used for the flag, the program will display the following usage message:

Usage: go run . [OPTION]

### _*Example*_

```bash
cat file.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $

go run . --reverse=file.txt
hello

```

## _*Instructions*_

- Project must be written in Go.
- Code must respect good practices.
- It is recommended to have test files for unit testing.

## _*Learning Objectives*_

- Understanding the Go file system (fs) API
- Data manipulation

---

*Author: [Raymond Ogwel](https://github.com/anxielray) 🚀
