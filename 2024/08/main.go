package main

import (
	"bufio"
	"fmt"
	"os"
)

type Vector struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	antMap := [][]byte{}
	for scanner.Scan() {
		antMap = append(antMap, scanner.Bytes())
	}

	getValidAntiNodes(antMap)
	getValidHarmonicAntiNodes(antMap)
}

func getValidAntiNodes(antMap [][]byte) {
	antennas := getAntennas(antMap)
	antiNodeLocations := make(map[int][]int)
	for _, coords := range antennas {
		for _, mainCoord := range coords {
			for _, floatingCoord := range coords {
				if mainCoord == floatingCoord {
					continue
				}

				dist := Vector{
					x: floatingCoord.x - mainCoord.x,
					y: floatingCoord.y - mainCoord.y,
				}

				if mainCoord.x-dist.x >= 0 &&
					mainCoord.y-dist.y >= 0 {
					antiNodeLocations[mainCoord.x-dist.x] =
						append(antiNodeLocations[mainCoord.x-dist.x], mainCoord.y-dist.y)
				}

				if floatingCoord.x+dist.x < len(antMap[0]) &&
					floatingCoord.y+dist.y < len(antMap) {
					antiNodeLocations[floatingCoord.x+dist.x] =
						append(antiNodeLocations[floatingCoord.x+dist.x], floatingCoord.y+dist.y)
				}
			}
		}
	}

	setMap := make(map[int]map[int]struct{}, 0)
	for x, ys := range antiNodeLocations {
		if x < 0 || x >= len(antMap[0]) {
			continue
		}

		for _, y := range ys {
			if y < 0 || y >= len(antMap) {
				continue
			}
			if setMap[x] == nil {
				setMap[x] = make(map[int]struct{})
			}
			setMap[x][y] = struct{}{}
		}
	}

	counter := 0
	for _, y := range setMap {
		counter += len(y)
	}
	fmt.Printf("unique AntiNodeSpots: %v\n", counter)
}

func getValidHarmonicAntiNodes(antMap [][]byte) {
	antennas := getAntennas(antMap)
	antiNodeLocations := make(map[int][]int)
	for _, coords := range antennas {
		for _, mainCoord := range coords {
			for _, floatingCoord := range coords {
				if mainCoord == floatingCoord {
					continue
				}

				dist := Vector{
					x: floatingCoord.x - mainCoord.x,
					y: floatingCoord.y - mainCoord.y,
				}

				i := 0
				for ok := true; ok; ok = mainCoord.x-dist.x*i >= 0 &&
					mainCoord.x-dist.x*i < len(antMap[0]) &&
					mainCoord.y-dist.y*i >= 0 &&
					mainCoord.y-dist.y*i < len(antMap) {
					antiNodeLocations[mainCoord.x-dist.x*i] =
						append(antiNodeLocations[mainCoord.x-dist.x*i], mainCoord.y-dist.y*i)
					i++
				}

				i = 0
				for ok := true; ok; ok = floatingCoord.x+dist.x*i < len(antMap[0]) &&
					floatingCoord.x+dist.x*i >= 0 &&
					floatingCoord.y+dist.y*i < len(antMap) &&
					floatingCoord.y+dist.y*i >= 0 {
					antiNodeLocations[floatingCoord.x+dist.x*i] =
						append(antiNodeLocations[floatingCoord.x+dist.x*i], floatingCoord.y+dist.y*i)
					i++
				}
			}
		}
	}

	setMap := make(map[int]map[int]struct{}, 0)
	for x, ys := range antiNodeLocations {
		if x < 0 || x >= len(antMap[0]) {
			continue
		}

		for _, y := range ys {
			if y < 0 || y >= len(antMap) {
				continue
			}
			if setMap[x] == nil {
				setMap[x] = make(map[int]struct{})
			}
			setMap[x][y] = struct{}{}
		}
	}

	counter := 0
	for _, y := range setMap {
		counter += len(y)
	}
	fmt.Printf("unique harmonic AntiNodeSpots: %v\n", counter)
}

func getAntennas(antMap [][]byte) map[byte][]Vector {
	antennas := make(map[byte][]Vector)
	for y := range len(antMap) {
		for x := range len(antMap[0]) {
			if antMap[y][x] != '.' {
				antennas[antMap[y][x]] = append(antennas[antMap[y][x]], Vector{x: x, y: y})
			}
		}
	}
	return antennas
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
