package student

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFiles(str []string, name string) []string {
	filename := name + ".txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return str
}

func WriteFile(name string, str [][]string) {
	f, err := os.Open(name)
	if err != nil {
		file, err := os.Create(name)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()
		for index := 0; index < len(str); index++ {
			for index1 := 0; index1 < 8; index1++ {
				_, err2 := file.WriteString(str[index][index1])
				_, err2 = file.WriteString("\n")
				if err2 != nil {
					log.Fatal(err2)
				}
			}
		}
	} else {
		fmt.Printf("File with the name \"%s\" already exists!\nPlease delete that file or choose different name for your file.\n", name)
		os.Exit(0)
	}
	defer f.Close()
}

func StrByLines(arr [][8]string, length int, name string) []string {
	var array []string
	for index1 := 0; index1 < 8; index1++ {
		str := ""
		for index2 := 0; index2 < length; index2++ {
			for _, val := range arr[index2][index1] {
				if val != '\n' && index2 != length-1 {
					str = str + string(val)
				} else if index2 == length-1 {
					str = str + string(val)
				}
			}
		}
		array = append(array, str)
	}
	return array
}

func CheckThirdArg(str string) (string, bool) {
	res := "Not valid flag"
	if len(str) > 9 && (str[0:9] == "--output=") {
		return str[9:], true
	}
	return res, false
}

func Ascii(number int, str []string, index int) string {
	start := number*9 + 2 + index - 1
	nothing := "Some Problems Here"
	for index, val := range str {
		if index == start {
			return val
		}
	}
	return nothing
}
