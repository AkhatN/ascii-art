package ascii

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

//CheckValid alphabets
func CheckValid(a string) error {
	for _, char := range a {
		if char == 10 || char == 13 {
			continue
		}

		if char < 32 || char > 126 {
			return errors.New("Not Valid Characters")
		}
	}

	return nil
}

//CheckFile adsada
func CheckFile(a string) error {
	if a != "shadow" && a != "standard" && a != "thinkertoy" {
		return errors.New("Invalid banner")
	}

	return nil
}

//PrintAscii prints words in ascii
func PrintAscii(a, banner string) (string, error) {
	s := ""
	for _, char := range a {
		if char != rune(13) {
			s += string(char)
		}
	}

	banner = "ascii/" + banner + ".txt"
	file, err := os.Open(banner)
	defer file.Close()
	if err != nil {
		return "", errors.New("Banner not found")
	}

	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if len(data) != 855 {
		return "", errors.New("Banner is corrupted")
	}

	splittedWord := strings.Split(s, string(rune(10)))
	arrWord := ""
	answer := ""

	for _, word := range splittedWord {
		if word == "" {
			answer += "\n"
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					arrWord += data[int(char-32)*9+i]
					if len(word)-1 == j {
						arrWord += "\n"
					}
				}
			}
			answer += arrWord
			arrWord = ""
		}
	}
	return answer, nil
}
