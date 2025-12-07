package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ruleset, updates := extractInput("input")

	checkForCorrectUpdates(ruleset, updates)
	fixIncorrectUpdates(ruleset, updates)
}

func extractInput(filename string) ([][]int, [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	ruleset := [][]int{}
	scanner := bufio.NewScanner(file)
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			rule := strings.Split(line, "|")
			first, err := strconv.Atoi(rule[0])
			if err != nil {
				panic(err)
			}
			sec, err := strconv.Atoi(rule[1])
			if err != nil {
				panic(err)
			}

			ruleset = append(ruleset, []int{first, sec})
			continue
		}

		if !strings.Contains(line, ",") {
			continue
		}

		items := strings.Split(line, ",")
		intLine := []int{}
		for _, item := range items {
			intItem, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}

			intLine = append(intLine, intItem)
		}
		updates = append(updates, intLine)
	}

	return ruleset, updates
}

func checkForCorrectUpdates(ruleset, updates [][]int) {
	sum := 0

	for _, update := range updates {
		if ok, _ := correctUpdate(ruleset, update); ok {
			sum += update[(len(update)-1)/2]
		}
	}

	fmt.Println("calc sum of correct updates:", sum)
}

func fixIncorrectUpdates(ruleset, updates [][]int) {
	sum := 0
	for i := range updates {
		for true {
			ok, fault := correctUpdate(ruleset, updates[i])
			if ok {
				break
			}

			if fault == nil || len(fault) != 2 {
				fmt.Println("invalid fault: ", fault)
				break
			}

			tmp := updates[i][fault[0]]
			updates[i][fault[0]] = updates[i][fault[1]]
			updates[i][fault[1]] = tmp

			if ok, _ = correctUpdate(ruleset, updates[i]); ok {
				sum += updates[i][(len(updates[i])-1)/2]
				break
			}
		}
	}
	fmt.Println("calc sum of corrected updates:", sum)
}

func correctUpdate(ruleset [][]int, updateList []int) (bool, []int) {
	for i := range updateList {
		for j := range updateList {
			if i == j {
				continue
			}

			if i < j && rulesetContains(ruleset, []int{updateList[j], updateList[i]}) ||
				i > j && rulesetContains(ruleset, []int{updateList[i], updateList[j]}) {
				return false, []int{i, j}
			}
		}
	}

	return true, nil
}

func rulesetContains(ruleset [][]int, sub []int) bool {
	if l := len(sub); l != 2 {
		fmt.Println("wrong length, got length of ", l, "want length of 2")
		return false
	}
	for i := range ruleset {
		if l := len(ruleset[i]); l != 2 {
			fmt.Println("wrong length, got length of ", l, "want length of 2")
			return false
		}

		if ruleset[i][0] == sub[0] && ruleset[i][1] == sub[1] {
			return true
		}
	}
	return false
}
