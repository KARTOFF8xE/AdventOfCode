package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	reports := strings.Split(input, "\n")

	calcSafeReports(reports)
	calcSafeReportsWithDampener(reports)
}

func calcSafeReports(reports []string) {
	safeReports := 0
	for _, report := range reports {
		report := strings.Split(report, " ")

		if len(report) == 0 {
			fmt.Println("no levels in report")
			return
		}

		levels := []int{}
		for _, lvl := range report {
			level, err := strconv.Atoi(lvl)
			if err != nil {
				fmt.Println("level not parsable, report:", report)
				return
			}
			levels = append(levels, level)
		}

		counter := 0
		for i := range len(levels) - 1 {
			if levels[i] < levels[i+1] && levels[i]-levels[i+1] >= -3 {
				counter++
			}
			if levels[i] > levels[i+1] && levels[i]-levels[i+1] <= 3 {
				counter--
			}
		}
		if math.Abs(float64(counter)) == float64(len(report)-1) {
			safeReports++
		}
	}

	fmt.Println("Safe reports:", safeReports)
}

func calcSafeReportsWithDampener(reports []string) {
	safeReports := 0
	for _, report := range reports {
		report := strings.Split(report, " ")

		if len(report) == 0 {
			fmt.Println("no levels in report")
			return
		}

		levels := []int{}
		for _, lvl := range report {
			level, err := strconv.Atoi(lvl)
			if err != nil {
				fmt.Println("level not parsable, report:", report)
				return
			}
			levels = append(levels, level)
		}

		for i := range levels {
			manLevels := removeElement(levels, i)
			counter := 0
			for j := range len(manLevels) - 1 {
				if manLevels[j] < manLevels[j+1] && manLevels[j]-manLevels[j+1] >= -3 {
					counter++
				}
				if manLevels[j] > manLevels[j+1] && manLevels[j]-manLevels[j+1] <= 3 {
					counter--
				}
			}
			if math.Abs(float64(counter)) == float64(len(manLevels)-1) {
				safeReports++
				break
			}
		}
	}

	fmt.Println("Safe reports with Dampener:", safeReports)
}

func removeElement(slice []int, index int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:index], newSlice[index+1:]...)
}
