package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a = "a10 10 20b 20 30c 30 dd"
	space := " "
	for strings.Contains(a, space) {
		spaceIndex := strings.Index(a, space)
		word := a[:spaceIndex]
		i, err := strconv.Atoi(word)
		if err == nil {
			fmt.Print(fmt.Sprint(i) + " ")
		}
		a = a[spaceIndex+1:]
		a = strings.Trim(a, space)
	}
	i, err := strconv.Atoi(a)
	if err == nil {
		fmt.Println(i)
	}
}
