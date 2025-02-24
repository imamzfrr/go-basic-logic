package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateOddMatrix(t *testing.T) {
	expected := [][]int{
		{1, 3, 5},
		{1, 3, 5},
		{1, 3, 5},
	}
	assert.Equal(t, expected, generateOddMatrix(3))
}

func TestGenerateDiagonalMatrix(t *testing.T) {
	expected := [][]int{
		{1, 0, 0},
		{0, 3, 0},
		{0, 0, 5},
	}
	assert.Equal(t, expected, generateDiagonalMatrix(3))
}

func TestGenerateReverseDiagonalMatrix(t *testing.T) {
	expected := [][]int{
		{0, 0, 5},
		{0, 3, 0},
		{1, 0, 0},
	}
	assert.Equal(t, expected, generateReverseDiagonalMatrix(3))
}

func TestSoal04(t *testing.T) {
	expected := [][]int{
		{1, 4, 7, 10, 13, 16, 19, 22, 25},
		{28, 31, 33, 35, 37, 39, 41, 43, 45},
		{48, 51, 53, 55, 57, 59, 61, 63, 65},
		{68, 71, 73, 75, 77, 79, 81, 83, 85},
		{88, 91, 93, 95, 97, 99, 101, 103, 105},
		{108, 111, 113, 115, 117, 119, 121, 123, 125},
		{128, 131, 133, 135, 137, 139, 141, 143, 145},
		{148, 151, 153, 155, 157, 159, 161, 163, 165},
		{168, 171, 173, 175, 177, 179, 181, 183, 185},
	}
	assert.Equal(t, expected, soal04(9))
}

func TestSoal05(t *testing.T) {
	expected := [][]int{
		{1, 3, 5, 7, 9, 11, 13, 15, 17},
		{35, 33, 31, 29, 27, 25, 23, 21, 19},
		{37, 39, 41, 43, 45, 47, 49, 51, 53},
		{71, 69, 67, 65, 63, 61, 59, 57, 55},
		{73, 75, 77, 79, 81, 83, 85, 87, 89},
		{107, 105, 103, 101, 99, 97, 95, 93, 91},
		{109, 111, 113, 115, 117, 119, 121, 123, 125},
		{143, 141, 139, 137, 135, 133, 131, 129, 127},
		{145, 147, 149, 151, 153, 155, 157, 159, 161},
	}
	assert.Equal(t, expected, soal05(9))
}

func TestSoal06(t *testing.T) {
	expected := [][]int{
		{1, 4, 7, 10, 13, 16, 19, 22, 25},
		{44, 42, 40, 38, 36, 34, 32, 30, 28},
		{46, 49, 52, 55, 58, 61, 64, 67, 70},
		{89, 87, 85, 83, 81, 79, 77, 75, 73},
		{91, 94, 97, 100, 103, 106, 109, 112, 115},
		{134, 132, 130, 128, 126, 124, 122, 120, 118},
		{136, 139, 142, 145, 148, 151, 154, 157, 160},
		{179, 177, 175, 173, 171, 169, 167, 165, 163},
		{181, 184, 187, 190, 193, 196, 199, 202, 205},
	}
	assert.Equal(t, expected, soal06(9))
}

func TestSoal10_2(t *testing.T) {
	expected := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{3, 5, 0, 0, 0, 0, 0, 0, 0},
		{7, 9, 11, 0, 0, 0, 0, 0, 0},
		{13, 15, 17, 19, 0, 0, 0, 0, 0},
		{21, 23, 25, 27, 29, 0, 0, 0, 0},
		{31, 33, 35, 37, 39, 41, 0, 0, 0},
		{43, 45, 47, 49, 51, 53, 55, 0, 0},
		{57, 59, 61, 63, 65, 67, 69, 71, 0},
		{73, 75, 77, 79, 81, 83, 85, 87, 89},
	}
	assert.Equal(t, expected, soal10_2(9))
}

func TestSoal0212(t *testing.T) {
	expected := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 3, 0, 0, 0, 0, 0, 3, 1},
		{1, 3, 5, 0, 0, 0, 5, 3, 1},
		{1, 3, 5, 7, 0, 7, 5, 3, 1},
		{1, 3, 5, 7, 9, 7, 5, 3, 1},
		{1, 3, 5, 7, 0, 7, 5, 3, 1},
		{1, 3, 5, 0, 0, 0, 5, 3, 1},
		{1, 3, 0, 0, 0, 0, 0, 3, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 1},
	}
	assert.Equal(t, expected, soal12(9))
}

// Logic 1
// /Soal 1 & 2
func TestGenerateSlice(t *testing.T) {
	expected := []int{1, 3, 5, 7, 9}
	assert.Equal(t, expected, generateSlice(1, 2, 5))
}

// /Soal 2-6
func TestGenerateReverseSlice(t *testing.T) {
	expected := []int{10, 8, 6, 4, 2}
	assert.Equal(t, expected, generateReverseSlice(10, 2, 5))
}

func TestSoal10(t *testing.T) {
	expected := []interface{}{1, "fizz", 2, "fizz", 4}
	assert.Equal(t, expected, soal10(5))
}

func TestSoal11(t *testing.T) {
	expected := []interface{}{"buzz", 1, "buzz", 2, "buzz"}
	assert.Equal(t, expected, soal11(5))
}

func TestGeneratePattern(t *testing.T) {
	assert.Equal(t, []int{1, 2, 2, 1}, GeneratePattern(4, 1, 1))
	assert.Equal(t, []int{1, 3, 5, 5, 3, 1}, GeneratePattern(6, 1, 2))
	assert.Equal(t, []int{2, 4, 6, 8, 6, 4, 2}, GeneratePattern(7, 2, 2))
}

func TestSoal0112(t *testing.T) {
	assert.Equal(t, []int{1, 3, 5, 7, 1, 3, 5, 7}, Soal12(8))
	assert.Equal(t, []int{1, 3, 5, 7, 1, 3, 5, 7, 1, 3}, Soal12(10))
}
