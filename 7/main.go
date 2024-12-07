package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	findValidSolutions(lines)
	findValidSolutionsWithConcat(lines)
}

func findValidSolutions(lines []string) {
	var sum int64 = 0
	for _, line := range lines {
		solution, err := strconv.ParseInt(strings.Split(line, ":")[0], 10, 64)
		if err != nil {
			panic(fmt.Sprint("Err1:", err))
		}

		elements := strings.Split(strings.Split(line, ":")[1], " ")
		elemInts := []int64{}
		for _, element := range elements {
			if element == "" {
				continue
			}

			elemInt, err := strconv.ParseInt(element, 10, 64)
			if err != nil {
				panic(fmt.Sprint("Err2:", err))
			}

			elemInts = append(elemInts, elemInt)
		}

		counter := int64(powInt(2, len(elemInts)-1))
		for i := 0; i <= powInt(2, len(elemInts)-1); i++ {
			sol := int64(elemInts[0])
			for j := 1; j < len(elemInts); j++ {
				if strconv.FormatInt(counter, 2)[j] == '0' {
					sol = sol + elemInts[j]
				} else {
					sol = sol * elemInts[j]
				}
			}
			if sol == solution {
				sum += solution
				break
			}
			counter++
		}
	}

	fmt.Println("sum of true equations (2 operands):", sum)
}

func findValidSolutionsWithConcat(lines []string) {
	var sum int64 = 0
	for _, line := range lines {
		solution, err := strconv.ParseInt(strings.Split(line, ":")[0], 10, 64)
		if err != nil {
			panic(err)
		}

		elements := strings.Split(strings.Split(line, ":")[1], " ")
		elemIntsFresh := []int64{}
		for _, element := range elements {
			if element == "" {
				continue
			}

			elemInt, err := strconv.ParseInt(element, 10, 64)
			if err != nil {
				panic(err)
			}

			elemIntsFresh = append(elemIntsFresh, elemInt)
		}

		counter := int64(powInt(3, len(elemIntsFresh)-1))
		for i := 0; i <= powInt(3, len(elemIntsFresh)-1); i++ {
			counterBytes := []byte(strconv.FormatInt(counter, 3))
			elemInts := make([]int64, len(elemIntsFresh))
			copy(elemInts, elemIntsFresh)
			sol := int64(elemInts[0])

			for j := 1; j < len(elemInts); j++ {
				switch counterBytes[j] {
				case '0':
					sol = sol + elemInts[j]
				case '1':
					sol = sol * elemInts[j]
				default:
					one := strconv.FormatInt(sol, 10)
					two := strconv.FormatInt(elemInts[j], 10)
					sol, err = strconv.ParseInt(fmt.Sprint(one+two), 10, 64)
					if err != nil {
						panic(err)
					}
				}
			}
			if sol == solution {
				sum += solution
				break
			}
			counter++
		}
	}

	fmt.Println("sum of true equations (3 operands):", sum)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func deleteElement[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
