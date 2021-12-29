package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"ascii-art-color/ascii"
)

func main() {
	arguments := os.Args[1:]
	if ascii.IsHelp(arguments) {
		return
	}

	if !ascii.CheckValidLen(arguments) {
		return
	}

	data, iscolor, ok := ascii.CheckColor(arguments)
	if !ok {
		return
	}

	var er bool
	var args *ascii.Args

	if !iscolor {
		if !ascii.IsLenOk(data.Words) {
			return
		}
	} else {
		args, er = ascii.IsValidArg(data.Separators)
	}

	if er {
		return
	}

	file, err := os.Open("standard.txt")
	defer file.Close()

	if err != nil {
		log.Println(err)
		return
	}

	var words, banner []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		banner = append(banner, scanner.Text())
	}

	if len(banner) != 855 {
		log.Println("File is corrupted")
		return
	}

	for _, word := range data.Words {
		argument := strings.Split(word, "\\n")
		for _, w := range argument {
			words = append(words, w)
		}
	}

	if !iscolor {
		ascii.PrintColor(words, banner, data.Color)
	} else {
		if args.Index == 0 {
			ascii.PrintColor(words, banner, data.Color)
		} else if args.Index == 1 {
			ascii.PrintSymbol(args.Arg1, words, banner, data.Color)
		} else if args.Index == 2 {
			ascii.PrintFromIndex(args.Arg1, words, banner, data.Color)
		} else if args.Index == 3 {
			ascii.PrintInRange(args.Arg1, args.Arg2, words, banner, data.Color)
		}
	}
}
