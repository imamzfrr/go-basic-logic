package main

import (
	"fmt"
	slice "github.com/imamzfrr/go-print-slice"
)

// Logic 1
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
		row := n - 1 - i           // Menentukan baris dari bawah ke atas
		col := i                   // Kolom berjalan normal
		matrix[row][col] = 1 + 2*i // Mengisi dengan bilangan ganjil
	}

	return matrix
}
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

//Logic 2

func soal04(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	num := 1 // Nilai awal di baris pertama
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

	for i := 0; i < n; i++ {
		value := 1
		for j := 0; j <= i; j += 3 {
			matrix[i][j] = value
			if j+1 < n {
				matrix[i][n-1-j] = 17
			}
			value += 2
		}
	}
	return matrix
}

func soal10(n int) []interface{} {
	slice := make([]interface{}, n)
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			slice[i-1] = i
		} else {
			slice[i-1] = "fizz"
		}
	}
	return slice
}

func soal11(n int) []interface{} {
	slice := make([]interface{}, n)

	for i := 0; i < n; i++ {
		if (i+1)%2 == 0 {
			slice[i] = "buzz"
		} else {
			if i < 2 {
				slice[i] = i + 1
			} else {
				if val, ok := slice[i-2].(int); ok {
					slice[i] = val * 2
				}
			}
		}
	}
	return slice
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

	fmt.Println("Soal 10:")
	slice.PrintSliceString(soal10(n))
	fmt.Println()

	fmt.Println("Soal 11:")
	slice.PrintSliceString(soal11(n))
	fmt.Println()

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

	fmt.Println("Soal 12:") //masih error
	slice.PrintSlice2(soal12(m))
}
