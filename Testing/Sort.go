package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

type User struct {
	name  string
	age   int
	score []int
}

func (u User) HiScore() int {
	hi := 0
	for _, sc := range u.score {
		if hi < sc {
			hi = sc
		}
	}
	return hi
}
func main() {
	user := User{"Poma", 23, []int{12, 15, 14, 56}}
	fmt.Println(user.age)
	fmt.Println(user.HiScore())
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}

func main() {
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, q := 0, len(num)-1; i < q; i, q = i+1, q-1 {
		num[i], num[q] = num[q], num[i]
	}
	fmt.Println(num)
}

func main() {
	num := []int{1, 3, 2, 4, 5, 9, 7, 8, 6, 10}
	a := sort.Reverse(sort.IntSlice(num))
	fmt.Println(a)

	numrev := []int{}
	for i := len(num) - 1; i >= 0; i-- {
		numrev = append(numrev, num[i])
	}
	fmt.Println(numrev)
}

func run() int {
	time.Sleep(time.Second * 1)
	return 10
}

func init() {
	log.SetFormatter()
}
func main() {
	for i := 0; i < 3; i++ {
		a := run()
		log.Println("hi", a)
		log.Info("result")
	}
}
