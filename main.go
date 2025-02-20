package main

import (
	"fmt"
	slice "github.com/imamzfrr/go-print-slice"
)

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

func generateHighlightedMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			// Kosongkan semua sel terlebih dahulu
			matrix[i][j] = 0
		}
	}

	for i := 0; i < n; i++ {
		matrix[i][i] = 1 + 2*i // Diagonal utama
		if i%2 == 0 && i < n-1 {
			matrix[i][i+1] = 1 + 2*(i+1) // Pola tambahan ke kanan
		}
		if i%2 == 1 && i > 0 {
			matrix[i][i-1] = 1 + 2*(i-1) // Pola tambahan ke kiri
		}
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

func soal10(n int) []int {
	slice := make([]int, n)
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			slice[i-1] = -1 // Menandai 'fizz' dengan -1
		} else {
			slice[i-1] = i
		}
	}
	return slice
}

func soal11(n int) []int {
	slice := make([]int, n)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			slice[i-1] = -2 // Menandai 'buzz' dengan -2
		} else {
			slice[i-1] = i
		}
	}
	return slice
}

func main() {
	n := 10

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
	slice.PrintSlice(soal10(n))
	fmt.Println()

	fmt.Println("Soal 11:")
	slice.PrintSlice(soal11(n))
	fmt.Println()

	m := 9
	fmt.Println("Soal 01:")
	slice.PrintSlice2(generateOddMatrix(m))

	fmt.Println("Soal 07:")
	slice.PrintSlice2(generateDiagonalMatrix(m))

	fmt.Println("Soal 13:")
	slice.PrintSlice2(generateHighlightedMatrix(m))
}
