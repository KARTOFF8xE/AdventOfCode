package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := getInput()
	lines := strings.Split(input, "\n")

	col0, col1 := getCols(lines)
	sort.Ints(col0)
	sort.Ints(col1)

	calcDist(col0, col1)
	calcSimScore(col0, col1)
}

func getCols(lines []string) ([]int, []int) {
	col0 := []int{}
	col1 := []int{}
	for _, line := range lines {
		first, err := strconv.Atoi(line[0:strings.Index(line, " ")])
		if err != nil {
			fmt.Println(err)
		}
		col0 = append(col0, first)

		second, err := strconv.Atoi(strings.TrimSpace(line[strings.Index(line, " "):]))
		if err != nil {
			fmt.Println(err)
		}
		col1 = append(col1, second)
	}
	return col0, col1
}

func calcDist(col0, col1 []int) {
	dist := 0
	for i := range col0 {
		if len(col1) < i {
			fmt.Println("col0 larger col1")
			return
		}
		dist += int(math.Abs(float64(col0[i] - col1[i])))
	}
	fmt.Println("Distance:", dist)
}

func calcSimScore(col0, col1 []int) {
	if len(col0) != len(col1) {
		fmt.Println("lengths not equal")
		return
	}

	simScore := 0
	for _, val := range col0 {
		for i := range col1 {
			if val == col1[i] {
				simScore += val
			}
		}
	}
	fmt.Println("SimScore:", simScore)
}
