package main

import (
	"fmt"
	"os"
	"strings"
	"student"
)

func main() {
	arguments := os.Args[1:]
	sozder := strings.Split(arguments[0], "\\n")
	var str string
	var third bool
	var filearray [][]string
	for index := 0; index < len(sozder); index++ {
		var standard []string
		if len(sozder[index]) > 0 && len(arguments) == 3 {
			standard = student.ReadFiles(standard, arguments[1])
			str, third = student.CheckThirdArg(arguments[2])
			if !third {
				fmt.Println(str)
				return
			}
		} else {
			return
		}
		result := make([][8]string, len(sozder[index]))
		for index1, val := range sozder[index] {
			for index2 := 0; index2 < 8; index2++ {
				result[index1][index2] = student.Ascii(int(val-32), standard, index2)
			}
		}
		filearray = append(filearray, student.StrByLines(result, len(sozder[index]), str))
	}
	student.WriteFile(str, filearray)
	return
}
