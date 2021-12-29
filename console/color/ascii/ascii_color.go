package ascii

import (
	"fmt"
	"strconv"
	"strings"
)

//Reset ...
var Reset = "\033[0m"

//Colors ...
var Colors = map[string]string{
	"reset":  "\033[0m",
	"red":    "\033[31m",
	"green":  "\033[32m",
	"yellow": "\033[33m",
	"blue":   "\033[34m",
	"purple": "\033[35m",
	"cyan":   "\033[36m",
	"white":  "\033[97m",
	"orange": "\033[38;2;255;128;0m",
}

//IsHelp ...
func IsHelp(args []string) bool {
	for _, word := range args {
		if word == "--help" {
			fmt.Println("1. Try: [word] - it will print the word with default color.")
			fmt.Println("Example: go run . hello")
			fmt.Println()
			fmt.Println("2. Try: [word] [--color=<color>] - it will print the word with the specified color.")
			fmt.Println("Example: go run . hello --color=green")
			fmt.Println()
			fmt.Println("3. Try: [word] [--color=<color>] [number] - it will print the exact symbol of the word with the specified color.")
			fmt.Println("Example: go run . hello --color=red 2")
			fmt.Println()
			fmt.Println("4. Try: [word] [--color=<color>] [number] [:] - it will print the word from specified symbol until the end with the specified color.")
			fmt.Println("Example: go run . hello --color=purple 3 :")
			fmt.Println()
			fmt.Println("5. Try: [word] [--color=<color>] [number] [:] [number] - it will print the word in defined range of symbols of the word with the specified color.")
			fmt.Println("Example: go run . hello --color=yellow 2 : 3")
			return true
		}
	}

	return false
}

//CheckValidLen ...
func CheckValidLen(args []string) bool {
	if len(args) == 0 {
		fmt.Println("No arguments\nTry '--help' for more information.")
		return false
	}

	if len(args) > 5 {
		fmt.Println("Too many arguments")
		return false
	}

	for _, word := range args {
		for _, char := range word {
			if char < 32 || char > 126 { // Valid Characters
				fmt.Println("Invalid arguments\nAccepts symbols only from ascii table.")
				return false
			}
		}
	}

	return true
}

//Data ...
type Data struct {
	Color      string
	Words      []string
	Separators []string
}

//CheckColor ...
func CheckColor(s []string) (*Data, bool, bool) {
	var data Data
	data.Color = "\033[0m"
	str := ""
	sep := 0
	iscolor := false

	for i, col := range s {
		if strings.HasPrefix(col, "--color") && i == 1 {
			iscolor = true
			str = col

			if str == "--color" {
				fmt.Println("color: requires '=' after the flag\nTry '--help' for more information.")
				return &data, true, false
			}

			if str[7] == '=' {
				data.Color = str[8:]
			} else {
				fmt.Println("color: invalid flag\nTry '--help' for more information.")
				return &data, true, false
			}

			if data.Color == "" {
				data.Color = "\033[0m"

			} else {
				if val, ok := Colors[data.Color]; ok {
					data.Color = val

				} else {
					fmt.Println("There is no such color")
					fmt.Println("Try: red, green, yellow, blue, purple, cyan, white or orange.")
					return &data, true, false
				}
			}

			sep = i
			break
		}
		data.Words = append(data.Words, col)
	}
	if sep != len(s)-1 {
		for k := sep + 1; k < len(s); k++ {
			data.Separators = append(data.Separators, s[k])
		}
	}

	return &data, iscolor, true
}

//IsLenOk ...
func IsLenOk(s []string) bool {
	if len(s) != 1 {
		fmt.Println("There shoud be only one argument or second argument should be the flag '--color'\nTry '--help' for more information.")
		return false
	}

	return true
}

//Args ...
type Args struct {
	Arg1  int
	Arg2  int
	Index int
}

//IsValidArg ...
func IsValidArg(s []string) (*Args, bool) {
	var arg Args
	if len(s) == 0 {
		return &arg, false
	}

	a1, _ := strconv.Atoi(s[0])

	if !isNumber(a1, s[0]) {
		fmt.Printf("color: invalid index - '%v'\nTry '--help' for more information.\n", s[0])
		return &arg, true
	}
	if len(s) == 1 {
		arg.Arg1 = a1
		arg.Index = 1

	} else if len(s) == 2 {
		if s[1] != ":" {
			fmt.Printf("color: invalid index - '%v'\nTry '--help' for more information.\n", s[1])
			return &arg, true
		}

		arg.Arg1 = a1
		arg.Index = 2
	} else {
		if s[1] != ":" {
			fmt.Printf("color: invalid index - '%v'\nTry '--help' for more information.\n", s[1])
			return &arg, true
		}

		a2, _ := strconv.Atoi(s[2])
		if !isNumber(a2, s[2]) {
			fmt.Printf("color: invalid index - '%v'\nTry '--help' for more information.\n", s[2])
			return &arg, true
		}

		arg.Arg1 = a1
		arg.Arg2 = a2
		arg.Index = 3
	}
	return &arg, false
}

//isNumber ...
func isNumber(n int, s string) bool {
	if n == 0 && (s != "0" && s != "+0" && s != "-0") {
		return false
	}

	return true
}

//PrintColor ...
func PrintColor(words []string, data []string, col string) {
	arrWord := ""
	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					arrWord = data[int(char-32)*9+i]
					fmt.Print(col, arrWord, Reset)

					if len(word)-1 == j {
						arrWord = "\n"
						fmt.Print(Reset, arrWord)
					}
				}
			}
			arrWord = ""
		}
	}
}

func toPositive(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

//PrintSymbol ....
func PrintSymbol(n int, words []string, data []string, col string) {
	n = toPositive(n)
	isColor := true
	if n == 0 {
		isColor = false
	}

	if n != 0 {
		n--
	}

	arrWord := ""

	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					arrWord = data[int(char-32)*9+i]

					if j == n && isColor {
						fmt.Print(col, arrWord, Reset)
					} else {
						fmt.Print(arrWord)
					}

					if len(word)-1 == j {
						arrWord = "\n"
						fmt.Print(Reset, arrWord)
					}
				}
			}
			arrWord = ""
		}
	}
}

//PrintFromIndex ...
func PrintFromIndex(n int, words []string, data []string, col string) {
	n = toPositive(n)
	isColor := false

	if n != 0 {
		n--
	}

	arrWord := ""

	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					arrWord = data[int(char-32)*9+i]
					isColor = false

					if j >= n {
						isColor = true
					}

					if isColor {
						fmt.Print(col, arrWord, Reset)
					} else {
						fmt.Print(arrWord)
					}

					if len(word)-1 == j {
						arrWord = "\n"
						fmt.Print(Reset, arrWord)
					}
				}
			}
			arrWord = ""
		}
	}
}

//PrintInRange ...
func PrintInRange(n1, n2 int, words []string, data []string, col string) {
	n1 = toPositive(n1)
	n2 = toPositive(n2)
	m := 1

	if n1 > n2 {
		n1, n2 = n2, n1
	}

	if n2 == 0 {
		m = 0
	}

	if n1 != 0 {
		n1--
	}

	if n2 != 0 {
		n2--
	}

	isColor := false
	arrWord := ""

	for _, word := range words {
		if word == "" {
			fmt.Println()
		} else {
			for i := 1; i < 9; i++ {
				for j, char := range word {
					arrWord = data[int(char-32)*9+i]

					if j >= n1 && j <= n2 && m != 0 {
						isColor = true
					} else {
						isColor = false
					}

					if isColor {
						fmt.Print(col, arrWord, Reset)
					} else {
						fmt.Print(arrWord)
					}

					if len(word)-1 == j {
						arrWord = "\n"
						fmt.Print(Reset, arrWord)
					}
				}
			}
			arrWord = ""
		}
	}
}
