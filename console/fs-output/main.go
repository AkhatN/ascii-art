package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) > 3 {
		fmt.Println("Too many arguments!")
		os.Exit(0)
	} else { // Valid Characters
		for _, word := range args {
			for _, char := range word {
				if char < 32 || char > 126 {
					fmt.Println("NV")
					return
				}
			}
		}
		words, nameFont, newFile, output := wordsFontNewFileSplitted(args)
		file, _ := os.Open(nameFont)
		scanner := bufio.NewScanner(file)
		var splittedWords, data, arrWords []string
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		for _, word := range words {
			splittedWords = strings.Split(word, "\\n")
			for _, ar := range splittedWords {
				arrWords = append(arrWords, ar)
			}
		}
		file.Close()
		if output == false { //Case 1: Print
			arrWord := ""
			for _, word := range arrWords {
				if word == "" {
					fmt.Println()
				} else {
					for i := 1; i < 9; i++ {
						for j, char := range word {
							arrWord += data[int(char-32)*9+i]
							if len(word)-1 == j {
								arrWord += "\n"
							}
						}
					}
					arrWord = arrWord[:len(arrWord)-1]
					fmt.Println(arrWord)
					arrWord = ""
				}
			}
		} else { //Case 2: Write File
			newfile, er := os.OpenFile(newFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if er != nil {
				fmt.Println("Output needs a name of file")
				os.Exit(0)
			}
			arrWord := ""
			for _, word := range arrWords {
				if word == "" {
					newfile.Write([]byte("\n"))
				} else {
					for i := 1; i < 9; i++ {
						for j, char := range word {
							arrWord += data[int(char-32)*9+i]
							if len(word)-1 == j {
								arrWord += "\n"
							}
						}
					}
					newfile.Write([]byte(arrWord))
					arrWord = ""
				}
			}
			newfile.Close()
		}
	}

}
func wordsFontNewFileSplitted(x []string) ([]string, string, string, bool) {
	nameFont := "standard.txt"
	newFile := ""
	words := make([]string, 0)
	output := false
	for i := 0; i < len(x); i++ {
		if x[i] == "standard" || x[i] == "thinkertoy" || x[i] == "shadow" {
			nameFont = strings.ToLower(x[i]) + ".txt"
			continue
		}
		if strings.HasPrefix(x[i], "--output") {
			newFile = x[i]
			if newFile == "--output" {
				newFile = ""
				output = true
			} else if newFile == "--output=" {
				newFile = ""
				output = true
			} else if newFile[8] == '=' {
				newFile = newFile[9:]
				output = true
			} else {
				output = false
				words = append(words, x[i])
			}
			continue
		}
		words = append(words, x[i])
	}
	return words, nameFont, newFile, output
}
