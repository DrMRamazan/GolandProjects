package main

import "fmt"

type TestStruct struct {
	On    bool
	Ammo  int
	Power int
}

func main() {

	testStruct := &TestStruct{false, 1, 2}
	fmt.Println(testStruct.RideBike())
	fmt.Println(testStruct.Shoot())

}

func (counts *TestStruct) RideBike() bool {
	if !counts.On || counts.Power <= 0 {
		return false
	} else {
		counts.Power--
		return true
	}
}

func (counts *TestStruct) Shoot() bool {
	if !counts.On || counts.Ammo > 0 {
		return false
	} else {
		counts.Ammo--
		return true
	}
}
