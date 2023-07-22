package main

import "fmt"

// const r1 = 3
// const c1 = 5
// const r2 = 5
// const c2 = 4

// func main() {
// 	matrix1 := [r1][c1]int{
// 		{2, 4, 8, 3, 5},
// 		{3, 5, 7, 6, 9},
// 		{3, 1, 1, 5, 8},
// 	}

// 	matrix2 := [r2][c2]int{
// 		{5, 4, 6, 3},
// 		{2, 1, 1, 4},
// 		{3, 1, 4, 7},
// 		{3, 2, 5, 8},
// 		{3, 5, 4, 6},
// 	}

// 	result := mulmatrix(matrix1, matrix2)
// 	fmt.Println(result)
// }

// func mulmatrix(matrix1 [r1][c1]int, matrix2 [r2][c2]int) (result [r1][c2]int) {
// 	for i := 0; i < r1; i++ {
// 		for j := 0; j < c2; j++ {
// 			for k := 0; k < c1; k++ {
// 				result[i][j] += matrix1[i][k] * matrix2[k][j]
// 			}
// 		}
// 	}
// 	return

// }

// const (
// 	row1 = 3
// 	col1 = 5
// 	row2 = 5
// 	col2 = 4
// )

// func main() {
// 	matrix1 := [row1][col1]int{
// 		// {5, 7, 9, 2, 11},
// 		// {2, 8, 1, 5, 8},
// 		// {6, 0, 1, 7, 1},
// 		{2, 4, 8, 3, 5},
// 		{3, 5, 7, 6, 9},
// 		{3, 1, 1, 5, 8},
// 	}

// 	matrix2 := [row2][col2]int{
// 		// {1, 5, 2, 8},
// 		// {8, 4, 2, 1},
// 		// {45, 12, 31, 2},
// 		// {7, 43, 14, 7},
// 		// {8, 9, 1, 7},
// 		{5, 4, 6, 3},
// 		{2, 1, 1, 4},
// 		{3, 1, 4, 7},
// 		{3, 2, 5, 8},
// 		{3, 5, 4, 6},
// 	}
// 	result := multiplication(matrix1, matrix2)
// 	fmt.Print(result)
// 	//print(result)

// }

// func multiplication(matrix1 [row1][col1]int, matrix2 [row2][col2]int) (result [row1][col2]int) {
// 	for i := 0; i < row1; i++ { /* каждая строка M1 */
// 		for j := 0; j < col2; j++ { /* каждый столбец M2 */
// 			for k := 0; k < col1; k++ {
// 				result[i][j] += matrix1[i][k] * matrix2[k][j]
// 			}
// 		}
// 	}
// 	return
// }

// func print(matrix [row1][col2]int) {
// 	for i := 0; i < row1; i++ { /* каждая строка M1 */
// 		for j := 0; j < col2; j++ { /* каждый столбец M2 */
// 			fmt.Print(matrix[i][j], " ")
// 		}
// 		fmt.Println()
// 	}
// }

const rows = 3
const cols = 4

func main() {
	matrix := [rows][cols]int{
		{5, 2, 3, 8},
		{3, 6, 8, 9},
		{5, 5, 7, 8},
	}

	fmt.Println(summatrix(matrix, matrix))
	res := summatrix(matrix, matrix)
	result(res)

	diagonalsum(res)
	fmt.Println(trans(res))
	one(matrix)
	deter(matrix)
}

func summatrix(a [rows][cols]int, b [rows][cols]int) [rows][cols]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
	return a
}

func result(m [rows][cols]int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println()
	}

}

func diagonalsum(a [rows][cols]int) {
	sum := 0

	for i := 0; i < len(a); i++ {
		sum += a[i][i]
	}
	fmt.Println(sum)
}

func trans(a [rows][cols]int) [cols][rows]int {
	var trans [cols][rows]int
	for i := 0; i < len(a[0]); i++ {
		for j := 0; j < len(a); j++ {
			trans[i][j] = a[j][i]
		}

	}
	return trans

}

func one(m [rows][cols]int) {
	for i := 0; i < rows*cols; i++ {
		row := i / cols
		col := i % cols
		fmt.Print(m[row][col], " ")
	}
}
func deter(a [rows][cols]int) {
	x := a[1][1]*a[2][2] - a[2][1]*a[1][2]
	y := a[1][0]*a[2][2] - a[2][0]*a[1][2]
	z := a[1][0]*a[2][1] - a[2][0]*a[1][1]

	deter := a[0][0]*x - a[0][1]*y + a[0][2]*z
	fmt.Println(deter)

}
