package main

import (
	"fmt"
	slice "github.com/imamzfrr/go-print-slice"
)

// Logic 2
func generateOddMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = 1 + 2*j
		}
	}
	return matrix
}

func generateDiagonalMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		if i < n {
			matrix[i][i] = 1 + 2*i
		}
	}
	return matrix
}
func generateReverseDiagonalMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		row := n - 1 - i
		col := i
		matrix[row][col] = 1 + 2*i
	}

	return matrix
}

func soal04(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	for rIndex, row := range matrix {
		for cIndex := range row {
			row[cIndex] = num
			if rIndex == 0 || cIndex == 0 || cIndex == n-1 {
				num += 3
			} else {
				num += 2
			}
		}
	}
	return matrix
}

func soal05(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	for rIndex, rows := range matrix {
		for cIndex := range rows {
			if rIndex%2 == 0 {
				rows[cIndex] = num
				num += 2
			} else {
				rows[n-1-cIndex] = num
				num += 2
			}
		}
	}
	return matrix

}
func soal06(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	for rIndex, rows := range matrix {
		for cIndex := range rows {
			if rIndex%2 == 0 {
				rows[cIndex] = num
				num += 3
			} else {
				rows[n-1-cIndex] = num
				num += 2
			}
		}
	}
	return matrix
}

func soal10_2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j] = num
			num += 2
		}
	}
	return matrix
}

func soal12(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for col := 0; col < (n+1)/2; col++ {
		num := 1 + (col * 2)
		for row := col; row < n-col; row++ {
			matrix[row][col] = num
			matrix[row][n-1-col] = num
		}
	}
	return matrix
}

func soal10(n int) []interface{} {
	result := make([]interface{}, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = 1 << uint(i/2)
		} else {
			result[i] = "fizz"
		}
	}

	return result
}

func soal11(n int) []interface{} {
	result := make([]interface{}, n)

	prev := 1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = "buzz"
		} else {
			result[i] = prev
			prev *= 2
		}
	}

	return result
}

// Logic 1
func generateSlice(start, step, count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = start + (i * step)
	}
	return slice
}

func generateReverseSlice(start, step, count int) []int {
	slice := make([]int, count)
	for i := 0; i < count; i++ {
		slice[i] = start - (i * step)
	}
	return slice
}

func GeneratePattern(n, start, step int) []int {
	result := make([]int, n)
	half := n / 2

	for i := 0; i < half; i++ {
		value := start + (i * step)
		result[i] = value
		result[n-1-i] = value
	}

	if n%2 != 0 {
		result[half] = start + (half * step)
	}

	return result
}

func Soal12(n int) []int {
	result := make([]int, n)
	x := 1
	y := 2

	for i := 0; i < n; i++ {
		result[i] = x + (i%4)*y
	}
	return result
}

// Logic 3
func soalLogic301(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	count := 1

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			// Isi dari kiri ke kanan
			for col := 0; col <= row; col++ {
				matrix[row][col] = count
				count += 2
			}
		} else {
			// Isi dari kanan ke kiri
			for col := row; col >= 0; col-- {
				matrix[row][col] = count
				count += 2
			}
		}
	}
	return matrix
}

func soalLogic302(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	count := 1

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			// Isi dari kiri ke kanan
			for col := row; col < n; col++ {
				matrix[row][col] = count
				count += 2
			}
		} else {
			// Isi dari kanan ke kiri
			for col := n - 1; col >= row; col-- {
				matrix[row][col] = count
				count += 2
			}
		}
	}
	return matrix
}

func main() {
	n := 10 // Logic 1

	fmt.Println("Soal 1:")
	slice.PrintSlice(generateSlice(1, 2, n))
	fmt.Println()

	fmt.Println("Soal 2:")
	slice.PrintSlice(generateSlice(2, 2, n))
	fmt.Println()

	fmt.Println("Soal 3:")
	slice.PrintSlice(generateSlice(3, 3, n))
	fmt.Println()

	fmt.Println("Soal 4:")
	slice.PrintSlice(generateReverseSlice(19, 2, 10))
	fmt.Println()

	fmt.Println("Soal 5:")
	slice.PrintSlice(generateReverseSlice(20, 2, 10))
	fmt.Println()

	fmt.Println("Soal 6:")
	slice.PrintSlice(generateReverseSlice(30, 3, 10))
	fmt.Println()

	fmt.Println("Soal 7:")
	slice.PrintSlice(GeneratePattern(10, 1, 2))
	fmt.Println()

	fmt.Println("Soal 8:")
	slice.PrintSlice(GeneratePattern(10, 2, 2))
	fmt.Println()

	fmt.Println("Soal 9:")
	slice.PrintSlice(GeneratePattern(10, 3, 3))
	fmt.Println()

	fmt.Println("Soal 10:")
	slice.PrintSliceString(soal10(n))

	fmt.Println("Soal 11:")
	slice.PrintSliceString(soal11(n))

	fmt.Println("Soal 12:")
	slice.PrintSlice(Soal12(12))
	fmt.Println()

	fmt.Println("=======LOGIC 2 =======")

	m := 9 //Logic 2
	fmt.Println("Soal 01:")
	slice.PrintSlice2(generateOddMatrix(m))

	fmt.Println("Soal 04:")
	slice.PrintSlice2(soal04(m))

	fmt.Println("Soal 05:")
	slice.PrintSlice2(soal05(m))

	fmt.Println("Soal 06:")
	slice.PrintSlice2(soal06(m))

	fmt.Println("Soal 07:")
	slice.PrintSlice2(generateDiagonalMatrix(m))

	fmt.Println("Soal 08:")
	slice.PrintSlice2(generateReverseDiagonalMatrix(m))

	fmt.Println("Soal 10:")
	slice.PrintSlice2(soal10_2(m))

	fmt.Println("Soal 12:")
	slice.PrintSlice2(soal12(m))

	fmt.Println()
	fmt.Println("=======LOGIC 3 =======")
	fmt.Println()

	fmt.Println("Soal 1: ")
	slice.PrintSlice2(soalLogic301(9))

	fmt.Println("Soal 2: ")
	slice.PrintSlice2(soalLogic302(9))

}
