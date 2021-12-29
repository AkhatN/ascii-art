package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	checkValidLen(args)
	words, banner, flag := checkWordBannerFlag(args)
	var dataFile, newWords []string
	file, err := os.Open(banner)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		dataFile = append(dataFile, scanner.Text())
	}
	if len(dataFile) != 855 {
		log.Println("File is corrupted")
		return
	}
	newWords = strings.Split(words, "\\n")
	arrWord := ""
	for _, word := range newWords {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					numberSpace := isCountSpace(word)
					if char == ' ' && flag == "justify" {
						amountWords := isCountWord(word, dataFile, i)
						b := flagLine(arrWord, flag, numberSpace, amountWords)
						for i := 0; i < b; i++ {
							arrWord += " "
						}
						if j == len(word)-1 {
							arrWord = arrWord[:len(arrWord)-1]
							arrWord += "\n"
						}
						continue
					}
					arrWord += dataFile[int(char-32)*9+i]
					if len(word)-1 == j {
						arrWord += "\n"
					}
				}
				if flag != "justify" {
					a := flagLine(arrWord, flag, -1, -1)
					for i := 0; i < a; i++ { // center left right
						fmt.Printf(" ")
					}
				}
				fmt.Printf(arrWord)
				arrWord = ""
			}
		}
	}
	file.Close()
}

func isCountWord(word string, dataFile []string, i int) int {
	arrWord := ""
	for _, char := range word {
		if char != ' ' {
			arrWord += dataFile[int(char-32)*9+i]
		}
	}
	a := len(arrWord)
	return a
}

func flagLine(s, f string, numberSpace, a int) int {
	width := 0
	if f == "center" {
		width = (sizeOfterm() - len(s)) / 2
	} else if f == "right" {
		width = sizeOfterm() - len(s)
	} else if f == "justify" {
		width = (sizeOfterm() - a) / numberSpace
	}
	return width
}

func isCountSpace(s string) int {
	a := 0
	for _, char := range s {
		if char == ' ' {
			a++
		}
	}
	return a
}

func checkValidLen(args []string) {
	if len(args) == 0 {
		fmt.Println("No arguments")
		os.Exit(0)
	}
	if len(args) > 3 {
		fmt.Println("Too many arguments")
		os.Exit(0)
	}
	for _, word := range args {
		for _, char := range word {
			if char < 32 || char > 126 {
				fmt.Println("NV")
				os.Exit(0)
			}
		}
	}
}

func sizeOfterm() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("agat")
		os.Exit(0)
	}
	temp := strings.Split(string(out[:len(out)-1]), " ")
	width, _ := strconv.Atoi(temp[1])
	return width
}

func checkWordBannerFlag(s []string) (string, string, string) {
	words := s[0]
	banner := "standard.txt"
	flag := "left"
	for i := 1; i < len(s); i++ {
		if strings.HasPrefix(s[i], "--align") {
			if i != len(s)-1 {
				fmt.Println("Wrong order")
				os.Exit(0)
			}
			if s[i] == "--align" || s[i] == "--align=" {
				fmt.Println("--align: needs an argument")
				os.Exit(0)
			}
			flag = s[i]
			if len(s[i]) > 7 && flag[7] != '=' {
				fmt.Println("Wrong operator")
				os.Exit(0)
			}
			flag = flag[8:]
			if flag != "center" && flag != "justify" && flag != "right" && flag != "left" {
				fmt.Println("--align: invalid value")
				os.Exit(0)
			}
			continue
		}
		banner = s[i] + ".txt"
	}
	return words, banner, flag
}
