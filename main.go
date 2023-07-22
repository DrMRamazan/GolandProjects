package main

import (
	"MODULE_NAME/calendar"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := calendar.Date{}
	err := date.SetYear(-2019)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(27)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Year())

}
