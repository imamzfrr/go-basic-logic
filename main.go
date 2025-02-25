package main

import (
	"awesomeProject1/utils"
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

func soal09(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row == col && row+col <= n-1 {
				matrix[row][col] = num
			} else if row == col && row+col >= n-1 {
				matrix[row][col] = num
			}
			num += 2
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

	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row >= col && row+col <= n-1 {
				matrix[row][col] = num
			} else if row <= col && row+col >= n-1 {
				matrix[row][col] = num
			}
			num += 2
		}
	}
	return matrix
}

func soal13(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row <= col && row+col <= n-1 {
				matrix[row][col] = num
				matrix[n-1-row][col] = num
			}
			num += 2
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

func soalLogic303(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 2

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			// Isi dari kiri ke kanan
			for col := 0; col < n-row; col++ {
				matrix[row][col] = num
				num += 2
			}
		} else {
			// Isi dari kanan ke kiri
			for col := n - 1 - row; col >= 0; col-- {
				matrix[row][col] = num
				num += 3
			}
		}
	}
	return matrix
}

func soalLogic304(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := n - 1; col >= n-row-1; col-- {
				matrix[row][col] = num
				num += 2
			}
		} else {
			for col := n - row - 1; col < n; col++ {
				matrix[row][col] = num
				num += 2
			}
		}
	}
	return matrix
}

func soalLogic305(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	count := 1
	mid := (n - 1) / 2

	for row := 0; row <= mid; row++ {
		if row%2 == 0 {
			for col := 0; col <= row; col++ {
				matrix[row][col] = count
				matrix[n-1-row][col] = count
				matrix[row][n-1-col] = count
				matrix[n-1-row][n-1-col] = count
				count += 2
			}
		} else {
			for col := row; col >= 0; col-- {
				matrix[row][col] = count
				matrix[n-1-row][col] = count
				matrix[row][n-1-col] = count
				matrix[n-1-row][n-1-col] = count
				count += 2
			}
		}
	}
	return matrix
}

func soalLogic306(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1

	for row := 0; row < n; row++ {
		for col := 0; col < n-row; col++ {
			if row <= col {
				if row%2 == 0 {
					matrix[row][col] = num
					matrix[n-1-row][col] = num
				} else {
					matrix[row][n-1-col] = num
					matrix[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return matrix
}

func soalLogic307(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	mid := (n - 1) / 2

	for row := 0; row <= n/2; row++ {
		for col := 0; col < n; col++ {
			if row+col >= mid && col-row <= mid {
				if row%2 == 1 {
					matrix[row][col] = num
					matrix[n-1-row][col] = num
				} else {
					matrix[row][col] = num
					matrix[n-row-1][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return matrix
}

func soalLogic308(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	mid := (n - 1) / 2

	for col := 0; col <= n/2; col++ {
		for row := 0; row < n; row++ {
			if row+col >= mid && row-col <= mid {
				if col%2 == 1 {
					matrix[row][col] = num
					matrix[row][n-1-col] = num
				} else {
					matrix[row][col] = num
					matrix[n-1-row][col] = num

					matrix[row][n-1-col] = num
					matrix[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return matrix
}

func soalLogic309(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for row := 0; row <= n/2; row++ {
		for col := 0; col <= n/2; col++ {
			if row+col >= n/2 {
				val := 1 + 2*((row+col)-n/2)
				matrix[row][col] = val
				matrix[n-1-row][col] = val
				matrix[row][n-1-col] = val
				matrix[n-1-row][n-1-col] = val
			}
		}
	}

	return matrix
}

func soalLogic310(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	mid := 4
	for r := 0; r <= n/2; r++ {
		for c := 0; c <= n/2; c++ {
			if r+c >= mid {
				val := 1 + 2*(n-1-r-c)
				matrix[r][c] = val
				matrix[n-1-r][c] = val
				matrix[r][n-1-c] = val
				matrix[n-1-r][n-1-c] = val
			}
		}
	}
	return matrix
}

func soalLogic311(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// First pattern: diagonal elements
	value := 1
	for i := 0; i < n/2; i++ {
		// Apply value to symmetric positions
		matrix[i][i] = value
		matrix[n-1-i][i] = value
		value += 2
	}

	// Reset value for second pattern
	value = 1

	// Second pattern: zigzag in right half based on row parity
	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			// Even rows: right to middle
			for col := n - 1; col >= n/2; col-- {
				if row+col >= 8 {
					matrix[row][col] = value
					matrix[n-1-row][col] = value
					value += 2
				}
			}
		} else {
			// Odd rows: middle to right
			for col := n / 2; col < n; col++ {
				if row+col >= 8 {
					matrix[row][col] = value
					matrix[n-1-row][col] = value
					value += 2
				}
			}
		}
	}

	return matrix
}

func soalLogic312(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	val := 1
	for i := 0; i < n/2; i++ {
		matrix[i][n-1-i] = val
		matrix[n-1-i][n-1-i] = val
		val += 2
	}

	val = 1
	for r := 0; r <= n/2; r++ {
		if r%2 == 0 {
			for c := 0; c <= n/2; c++ {
				if r >= c {
					matrix[r][c] = val
					matrix[n-1-r][c] = val
					val += 2
				}
			}
		} else {
			for c := n / 2; c >= 0; c-- {
				if r >= c {
					matrix[r][c] = val
					matrix[n-1-r][c] = val
					val += 2
				}
			}
		}
	}
	return matrix
}

func soalLogic313(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	mid := (n - 1) / 2

	for row := 0; row <= mid; row++ {
		for col := 0; col <= mid; col++ {
			if (row+col)%2 == 0 && row+col >= 4 {
				val := 2*col + 1
				matrix[row][col] = val
				matrix[n-1-row][col] = val
				matrix[row][n-1-col] = val
				matrix[n-1-row][n-1-col] = val
			}
		}
	}

	return matrix
}

func soalLogic314(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1

	for c := 0; c < n; c++ {
		if c%2 == 0 {
			for r := 0; r < n; r++ {
				matrix[r][c] = num
				num += 2
			}
		} else {
			for r := n - 1; r >= 0; r-- {
				matrix[r][c] = num
				num += 2
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

	fmt.Println("Soal 13: ")
	slice.PrintSlice2(soal13(9))

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

	fmt.Println("Soal 09: ")
	slice.PrintSlice2(soal09(9))

	fmt.Println("Soal 10:")
	slice.PrintSlice2(soal10_2(m))

	fmt.Println("Soal 12:")
	slice.PrintSlice2(soal12(m))

	fmt.Println()
	fmt.Println("=======LOGIC 3 =======")
	fmt.Println()

	fmt.Println("Soal 1: ")
	utils.PrintSlice2(soalLogic301(9))

	fmt.Println("Soal 2: ")
	utils.PrintSlice2(soalLogic302(9))

	fmt.Println("Soal 3: ")
	utils.PrintSlice2(soalLogic303(9))

	fmt.Println("Soal 4: ")
	utils.PrintSlice2(soalLogic304(9))

	fmt.Println("Soal 5: ")
	utils.PrintSlice2(soalLogic305(9))

	fmt.Println("Soal 6: ")
	utils.PrintSlice2(soalLogic306(9))

	fmt.Println("Soal 7: ")
	utils.PrintSlice2(soalLogic307(9))

	fmt.Println("Soal 8: ")
	utils.PrintSlice2(soalLogic308(9))

	fmt.Println("Soal 9: ")
	utils.PrintSlice2(soalLogic309(9))

	fmt.Println("Soal 10: ")
	utils.PrintSlice2(soalLogic310(9))

	fmt.Println("Soal 11: ")
	utils.PrintSlice2(soalLogic311(9))

	fmt.Println("Soal 12: ")
	utils.PrintSlice2(soalLogic312(9))

	fmt.Println("Soal 13: ")
	utils.PrintSlice2(soalLogic313(9))

	fmt.Println("Soal 14: ")
	utils.PrintSlice2(soalLogic314(9))

}
