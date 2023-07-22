package main

import (
	"fmt"
	"time"
)

func main() {

	//Получаем текущую дату и время
	now := time.Now() //2021-08-05 19:23:21.516443297 +0400 +04 m=+0.000080777

	// Форматируем вывод даты
	fmt.Println(now.Format("02-01-2006 15:04:05"))

	// парсит дату и время в строковом представлении
	firstTime, err := time.Parse("2006/01/02 15-04", "2020/05/15 17-45")

	if err != nil {
		panic(err)
	}
	fmt.Println(firstTime)

	firstTime1 := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	secondTime := time.Date(2020, time.May, 15, 16, 45, 12, 0, time.Local)

	// func (t Time) After(u Time) bool
	// true если позже
	fmt.Println(firstTime1.After(secondTime)) // true

	// func (t Time) Before(u Time) bool
	// true если раньше
	fmt.Println(firstTime1.Before(secondTime)) // false

	// func (t Time) Equal(u Time) bool
	// true если равны
	fmt.Println(firstTime1.Equal(secondTime)) // false

	now1 := time.Date(2020, time.May, 15, 17, 45, 12, 0, time.Local)

	// func (t Time) Add(d Duration) Time
	// изменяет дату в соответствии с параметром - "продолжительностью"
	future := now1.Add(time.Hour * 12) // перемещаемся на 12 часов вперед

	// func (t Time) AddDate(years int, months int, days int) Time
	// изменяет дату в соответствии с параметрами - количеством лет, месяцев и дней
	past := now1.AddDate(-1, -2, -3) // перемещаемся на 1 год, два месяца и 3 дня назад

	// func (t Time) Sub(u Time) Duration
	// вычисляет время, прошедшее между двумя датами
	fmt.Println(future.Sub(past)) // 10332h0m0s
}
