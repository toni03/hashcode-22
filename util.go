package main

import (
	"strconv"
)

func byteArrayToOneDimensionalIntArray(in []byte) []int {
	out := make([]int, 0)
	for _, v := range in {
		if isIgnorableElement(v, 1) {
			continue
		}

		element, err := strconv.Atoi(string(v))
		check(err)

		out = append(out, element)
	}
	return out
}

func byteArrayToTwoDimensionalIntArray(in []byte) [][]int {
	out := make([][]int, 0)
	aux := make([]int, 0)
	for _, v := range in {
		if isIgnorableElement(v, 2) {
			continue
		}

		if string(v) == "\n" {
			out = append(out, aux)
			aux = make([]int, 0)
		} else {
			element, err := strconv.Atoi(string(v))
			check(err)

			aux = append(aux, element)
		}
	}
	out = append(out, aux)

	return out
}

func isIgnorableElement(v byte, dimensions int) bool {
	if string(v) == " " {
		return true
	}

	if dimensions == 1 && string(v) == "\n" {
		return true
	}

	return false
}


func twoDSlice(n, m int) [][]int {
	table := make([][]int, n, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m, m)
	}
	return table
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Useful for indexing slices WITHOUT repeated elements
// [1, 2, 3, 4] => {1: 0, 2: 1, 3: 2, 4: 3}
func sliceToMap(s []int) map[int]int {
	index := make(map[int]int, len(s))
	for i, val := range s {
		index[val] = i
	}
	return index
}

// Useful for indexing slices WITH repeated elements
// [1, 2, 1, 3, 4, 4] => {1: [0, 2], 2: [1], 3: [3], 4: [4, 5]}
func sliceToIndex(s []int) map[int][]int {
	index := make(map[int][]int, len(s))
	for key := range index {
		index[key] = []int{}
	}
	for i, val := range s {
		index[val] = append(index[val], i)
	}
	return index
}

func parseInt(s string) int {
	duration, err := strconv.Atoi(s)
	check(err)
	return duration
}
