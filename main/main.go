package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var words string
var font string
var flag string
var data []string
var newSplitWord []string

func main() {
	if len(os.Args[1:]) == 0 {
		return
	} else if len(os.Args[1:]) > 3 {
		fmt.Println("Too many arguments!")
		return
	}
	arg := os.Args[1:]
	words, font, flag = wordFontFlagSplit(arg)
	ReadFile()
	AsciiArt(newSplitWord, data)
}

func ReadFile() {
	file, err := os.Open(font + ".txt")
	defer file.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	newSplitWord = strings.Split(words, "\\n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
}

func AsciiArt(words []string, font []string) {
	var fil *os.File
	var er error
	// fmt.Println(len(words))
	if flag != "" {
		fil, er = os.Create(flag + ".txt")
		defer fil.Close()
		if er != nil {
			log.Println(er.Error())
			return
		}
	}
	for i := 0; i < len(words); i++ {
		word := ""
		word = words[i]
		for index1 := 0; index1 < 8; index1++ {
			line := ""
			if word != "" {
				for index2 := 0; index2 < len(word); index2++ {
					line = line + font[int((word[index2]-32))*9+index1+1]
				}
			} else {
				if flag != "" {
					fil.Write([]byte("\n"))
				} else {
					fmt.Println(line)
				}
				break
			}
			if flag != "" {
				fil.Write([]byte(line))
				fil.Write([]byte("\n"))
			} else {
				fmt.Println(line)
			}
		}
	}
	return
}

func wordFontFlagSplit(s []string) (string, string, string) { //{"asdasd", "asdsad", "asdasd"}
	word, flagTemp, font, flag := s[0], "", "standard", ""
	for i := 1; i < len(s); i++ {
		if strings.HasPrefix(s[i], "--output") {
			if i != len(s)-1 {
				fmt.Println("Not correct order!")
				return "", "", ""
			}
			if s[i] == "--output" || s[i] == "--output=" {
				fmt.Println("--output: needs an argument!")
				return "", "", ""
			}
			flagTemp = s[i]
			if flagTemp[8] != '=' {
				fmt.Println("Not correct operator!")
				return "", "", ""
			}
			flagTemp = flagTemp[9:]
			flag = flagTemp
			continue
		}
		font = s[i]
	}
	return word, font, flag
}
