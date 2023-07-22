package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name string

	age int

	grade int
}

func main() {
	mapStudent := make(map[string]*Student, 0)
	fmt.Println("Введите имя, возраст и курс студента через пробел:")
	var in = bufio.NewReader(os.Stdin)

	for i, err := in.ReadString('\n'); err != io.EOF; i, err = in.ReadString('\n') {
		sliseStruct := strings.Fields(i)

		if len(sliseStruct) < 3 {
			fmt.Println("что-то пошло не так... Пожалуйста, попробуйте ещё раз.")
			continue
		}
		studentName := sliseStruct[0]
		studentAge, errAge := strconv.Atoi(sliseStruct[1])
		studentGrade, errGrade := strconv.Atoi(sliseStruct[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка при обработке возраста студента или его номера курса ! Пожалуйста, попробуйте снова...\n")
			continue
		}
		student := Student{studentName, studentAge, studentGrade}
		if err := student.get(mapStudent); err == false {
			student.put(mapStudent)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова...\n")
		}
	}
	fmt.Println("Вывожу всех введённых студентов:")
	for _, value := range mapStudent {
		fmt.Println(value.info())
	}
}
func (s Student) info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}
func (s *Student) put(m map[string]*Student) {
	m[s.name] = s
}
func (s Student) get(m map[string]*Student) bool {
	_, x := m[s.name]
	return x
}
