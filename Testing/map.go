package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	flag.StringVar(&str, "str", "", "big string")

	var subStr string
	flag.StringVar(&subStr, "subStr", "", "sub string")

	flag.Parse()

	if findSubstring(str, subStr) == true {
		fmt.Println("Подстрока subStr есть в строке str")
	} else {
		fmt.Println("Подстрока subStr отсутствует в строке str")
	}
}

func findSubstring(str string, subStr string) bool {
	r := []rune(str)
	sr := []rune(subStr)
	if len(sr) == 0 {
		return false
	}
str:
	for i, ru := range r {
		if ru == sr[0] {
			for j, x := range sr[1:] {
				if r[i+j+1] != x {
					continue str
				}
			}
			return true
		}
	}
	return false
}
