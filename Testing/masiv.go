package main

import (
	"fmt"
)

/*const size = 5

func maxi(input [size]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}

func mini(input [size]int) int {
	min := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}

func main() {
	number := [size]int{10, 22, 15, 45, 3}

	fmt.Println(number)
	fmt.Println(maxi(number))
	fmt.Println(mini(number))
}*/

/*func main() {
	s := "hello"
	for _, a := range s {
		fmt.Println(a)
		fmt.Printf("%c\n", a)
	}

	m := map[int]int{
		8: 42,
		2: 6,
		4: 9,
		5: 3}
	delete(m, 2)
	fmt.Println(m[4] - m[5])
} */
/*func f(v ...int) {
	res := 20
	for _, a := range v {
		res -= a
	}
	fmt.Println(res)
}
func main() {
	v := []int{8, 5, 3}
	f(v...)

	s := []int{1, 2, 4, 6, 8}
	s[2] = s[1]
	//s[3] = s[2] + s[0]
	fmt.Println(s)
	fmt.Println(s[4] % s[3])

}*/

/*func main() {
/*five := [5]int{1, 2, 3, 4, 5}
two := [2]int{6, 7}

var n [len(five) + len(two)]int
copy(n[:], five[:])
copy(n[len(five):], two[:])
fmt.Println(n)

masone := [5]int{1, 2, 3, 4, 5}
mastwo := [4]int{22, 33, 44, 55}
var masall [len(masone) + len(mastwo)]int
copy(masall[:], masone[:])
copy(masall[len(masone):], mastwo[:])
fmt.Println(len(masall), masall)
var all [len(masone) + len(mastwo)]int
copy(all[len(masone):], mastwo[:])
fmt.Println(all)*/

/*array := [7]int{22, 1, 44, -15, 45, -2, 3}
fmt.Println(array)

for i := 0; i < len(array); i++ {
	for j := i; j < len(array); j++ {
		if array[i] > array[j] {
			t := array[i]
			array[i] = array[j]
			array[j] = t
		}
	}
}
fmt.Println(array)

*/

/*s := [6]int{5, 2, 6, 3, 1, 4}
for i := 0; i < len(s); i++ {
	for j := i; j < len(s); j++ {
		if s[i] > s[i] {
			t := s[i]
			s[i] = s[j]
			s[j] = t
		}
	}
}
fmt.Println(s)
for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
	s[i], s[j] = s[j], s[i]
}
fmt.Println(s)

array := []string{"Test 1", "Test 2", "Test 3"}
arrReversed := []string{}

for i := len(array) - 1; i >= 0; i-- {
	arrReversed = append(arrReversed, array[i])
}

// Вывод: [Test 3 Test 2 Test 1]
fmt.Println(arrReversed)

a := []int{5, 3, 4, 7, 8, 9}
sort.Slice(a, func(i, j int) bool {
	return a[i] < a[j]
})

for _, v := range a {
	fmt.Println(v)
}*/

/*s := []int{5, 2, 6, 3, 1, 4}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for _, a := range s {
		fmt.Println(a)
	}
}*/

/*const dl = 3
const sh = 3

func main() {
	arr := [dl][sh]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			time.Sleep(2 * time.Second)
			fmt.Println(arr[i][j])
			log.Println(arr[i])
		}

	}
}*/

/*var arr [3][4]int
for i := 0; i < 3; i++ {
	for j := 0; j < 4; j++ {
		arr[i][j] = i + j
	}
}

fmt.Println(arr)

s := make([]string, 4)
a := make([]int, 3)
fmt.Println(s)
fmt.Printf("%#v\n", s)
fmt.Printf("%#v", a)

var slice []int
fmt.Printf("%#v", slice)*/

/*var text string
var width int
fmt.Scanf("%s %d", &text, &width)

res1 := []byte(text)
if len(res1) > 6 {
	if len(res1) <= width {
		fmt.Println(string(res1))
	} else {
		str := "..."
		res2 := string(res1[:width])
		res := res2 + str
		fmt.Println(res)
	}
} else {
	fmt.Println(string(res1))
}*/

/*var a, b int
fmt.Scan(&a)
fmt.Scan(&b)
if a > 1000 || b > 1000 {
	fmt.Println("No")

} else {
	c := int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

var a int
fmt.Scan(&a)
fmt.Printf("The next number for the number %v is %v.\n", a, a+1)
fmt.Printf("The previous number for the number %v is %v.\n", a, a-1)*/

/*func main() {
	trio := [3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			trio[i][j] = rand.Int()
		}
		fmt.Println(trio[i])
	}

}*/

/*func init() {
	rand.Seed(time.Now().UnixNano())
	//we are seeding the rand variable with present time
	//so that we would get different output each time
}
func main() {
	randMatrix := make([][]int, 3)
	// we have created a slice with length 3
	//which can hold type []int, these can be of different length
	for i := 0; i < 3; i++ {
		randMatrix[i] = make([]int, 3)
		// we are creating a slice which can hold type int
	}
	generate(randMatrix)
	fmt.Println(randMatrix)
}
func generate(randMatrix [][]int) {
	for i, innerArray := range randMatrix {
		for j := range innerArray {
			randMatrix[i][j] = rand.Intn(100)
			//looping over each element of array and assigning it a random variable
		}
	}
}*/

/*var a = [][]float64{
	{6, 11, 1},
	{3, 3, 18},
	{4, 0.67, 6.67},
}

var b = [][]float64{
	{10.96597, 35.10765, 96.72356},
	{2.35765, -84.11256, 0.87932},
	{18.24689, 22.13579, 1.11123},
}

var c = [][]float64{
	{1, 3, 3, 1},
	{1, 4, 6, 4},
	{1, 5, 10, 10},
	{1, 6, 15, 20},
}

func main() {
	fmt.Printf("%g, %g, %g\n", det(a), det(b), det(c))
}

func det(a [][]float64) float64 {
	if len(a) == 1 {
		return a[0][0]
	}
	sign, d := 1, float64(0)
	for i, x := range a[0] {
		d += float64(sign) * x * det(excludeColumn(a[1:], i))
		sign *= -1
	}
	return d
}

func excludeColumn(a [][]float64, i int) [][]float64 {
	b := make([][]float64, len(a))
	n := len(a[0]) - 1
	for j, row := range a {
		r := make([]float64, n)
		copy(r[:i], row[:i])
		copy(r[i:], row[i+1:])
		b[j] = r
	}
	return b
}
*/

const rows = 3
const cols = 3

func main() {
	matrix := [rows][cols]int{
		{1, 0, 1},
		{2, 3, 3},
		{4, 5, 5},
	}

	// for i := 0; i < len(matrix); i++ {
	// 	fmt.Println(matrix[i])
	// }

	// for i := 0; i < len(matrix); i++ {
	// 	for j := 0; j < len(matrix[i]); j++ {
	// 		fmt.Println(matrix[i][j])
	// 	}
	// }

	fmt.Println(findMaxInMatrix(matrix))

	fmt.Println(sumAll(matrix))
	fmt.Println(sumNotEven(matrix))
	fmt.Println(findMin(matrix))
}

func findMaxInMatrix(matrix [rows][cols]int) [rows]int {

	var result [rows]int
	for i := 0; i < len(matrix); i++ {
		lMax := maximum(matrix[i])
		result[i] = lMax

	}
	return result
}

func findMin(matrix [rows][cols]int) [rows]int {
	var result [rows]int
	for i := 0; i < len(matrix); i++ {
		lMin := minimum(matrix[i])
		result[i] = lMin
	}
	return result
}
func sumAll(matrix [rows][cols]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum += matrix[i][j]
		}
	}
	return sum
}

func sumNotEven(matrix [rows][cols]int) int {
	s := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if j%2 != 0 {
				s += matrix[i][j]

			}
		}
	}
	return s
}

func minimum(input [cols]int) int {
	min := input[0]

	for i := 0; i < len(input); i++ {
		if input[i] < min {
			min = input[i]
		}
	}
	return min
}
func maximum(input [cols]int) int {
	max := input[0]
	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
	}
	return max
}
