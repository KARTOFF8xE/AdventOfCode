package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	input := []string{}
	if scanner.Scan() {
		input = strings.Split(scanner.Text(), " ")
	}
	stones := []uint64{}
	for _, number := range input {
		stone, err := strconv.ParseInt(number, 10, 64)
		if err != nil {
			panic(err)
		}

		stones = append(stones, uint64(stone))
	}

	QuantityAfterBlinking(stones, 25)
	QuantityAfterBlinking(stones, 75)
}

func QuantityAfterBlinking(stones []uint64, leafeGeneration int) {
	counter := uint64(0)
	known := make(map[int]map[uint64]uint64)
	for _, stone := range stones {
		evolveStone(stone, 0, leafeGeneration, &counter, known)
	}
	fmt.Println("Stones after", leafeGeneration, "Generations:", counter)
}

func evolveStone(stone uint64, generation, leafeGeneration int, counter *uint64, known map[int]map[uint64]uint64) {
	if generation == leafeGeneration {
		*counter++
		return
	}
	if known[generation] != nil {
		if _, ok := known[generation][stone]; ok {
			*counter += known[generation][stone]
			return
		}
	}

	generation++
	if stone == 0 {
		oldCounter := *counter
		evolveStone(1, generation, leafeGeneration, counter, known)
		if known[generation-1] == nil {
			known[generation-1] = make(map[uint64]uint64)
		}
		known[generation-1][stone] = *counter - oldCounter
		return
	}

	oldCounter := *counter
	if len(strconv.FormatInt(int64(stone), 10))%2 != 0 {
		evolveStone(stone*2024, generation, leafeGeneration, counter, known)
	} else {
		stoneString := strconv.FormatInt(int64(stone), 10)
		left, err := strconv.ParseInt(stoneString[:len(stoneString)/2], 10, 64)
		if err != nil {
			panic(err)
		}
		right, err := strconv.ParseInt(stoneString[len(stoneString)/2:], 10, 64)
		if err != nil {
			panic(err)
		}
		evolveStone(uint64(left), generation, leafeGeneration, counter, known)
		evolveStone(uint64(right), generation, leafeGeneration, counter, known)
	}
	if known[generation-1] == nil {
		known[generation-1] = make(map[uint64]uint64)
	}
	known[generation-1][stone] = *counter - oldCounter
}
