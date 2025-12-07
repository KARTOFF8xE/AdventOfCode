package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	matrix := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matrix = append(matrix, []byte(scanner.Text()))
	}
	searchXmasAllDirections(matrix)
	searchXMAS(matrix)
}

func searchXmasAllDirections(matrix [][]byte) {
	filter0 := [][]byte{
		{88, 77, 65, 83},
	}
	filter1 := [][]byte{
		{83, 65, 77, 88},
	}
	filter2 := [][]byte{
		{83},
		{65},
		{77},
		{88},
	}
	filter3 := [][]byte{
		{88},
		{77},
		{65},
		{83},
	}
	filter4 := [][]byte{
		{83, 0, 0, 0},
		{0, 65, 0, 0},
		{0, 0, 77, 0},
		{0, 0, 0, 88},
	}
	filter5 := [][]byte{
		{88, 0, 0, 0},
		{0, 77, 0, 0},
		{0, 0, 65, 0},
		{0, 0, 0, 83},
	}
	filter6 := [][]byte{
		{0, 0, 0, 88},
		{0, 0, 77, 0},
		{0, 65, 0, 0},
		{83, 0, 0, 0},
	}
	filter7 := [][]byte{
		{0, 0, 0, 83},
		{0, 0, 65, 0},
		{0, 77, 0, 0},
		{88, 0, 0, 0},
	}
	counter := findPattern(matrix, filter0)
	counter += findPattern(matrix, filter1)
	counter += findPattern(matrix, filter2)
	counter += findPattern(matrix, filter3)
	counter += findPattern(matrix, filter4)
	counter += findPattern(matrix, filter5)
	counter += findPattern(matrix, filter6)
	counter += findPattern(matrix, filter7)
	fmt.Println("XMAS in all directions:", counter)
}

func searchXMAS(matrix [][]byte) {
	filter0 := [][]byte{
		{77, 0, 77},
		{0, 65, 0},
		{83, 0, 83},
	}
	filter1 := [][]byte{
		{77, 0, 83},
		{0, 65, 0},
		{77, 0, 83},
	}
	filter2 := [][]byte{
		{83, 0, 83},
		{0, 65, 0},
		{77, 0, 77},
	}
	filter3 := [][]byte{
		{83, 0, 77},
		{0, 65, 0},
		{83, 0, 77},
	}

	counter := findPattern(matrix, filter0)
	counter += findPattern(matrix, filter1)
	counter += findPattern(matrix, filter2)
	counter += findPattern(matrix, filter3)

	fmt.Println("X-Mas:", counter)
}

func findPattern(matrix, filter [][]byte) int {
	counter := 0
	matrixX := len(matrix[0])
	matrixY := len(matrix)
	filterX := len(filter[0])
	filterY := len(filter)

	match := 0
	for i := range filterY {
		for j := range filterX {
			match += int(filter[i][j]) * int(filter[i][j])
		}
	}

	for i := range matrixY - filterY + 1 {
		for j := range matrixX - filterX + 1 {
			test := 0
			for fi := range filterY {
				for fj := range filterX {
					test += int(matrix[i+fi][j+fj]) * int(filter[fi][fj])
				}
			}

			if test == match {
				counter++
			}
		}
	}

	return counter
}
