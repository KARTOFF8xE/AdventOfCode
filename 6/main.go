package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Vector struct {
	x int
	y int
}

var direction = map[Direction]Vector{
	North: Vector{x: 0, y: -1},
	East:  Vector{x: 1, y: 0},
	South: Vector{x: 0, y: 1},
	West:  Vector{x: -1, y: 0},
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	obstMap := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			obstMap = append(obstMap, line)
		}
	}

	predMap := predictPath(obstMap)
	setObstacles(predMap)
}

func setObstacles(obsMap []string) {
	possibilities := 0
	guardStart := locateGuard(obsMap)
	if guardStart.x < 0 || guardStart.y < 0 {
		panic("Guard not found")
	}
	for y := range obsMap {
		for x := range obsMap[0] {
			if obsMap[y][x] != 'X' {
				continue
			}
			if guardStart.x == x && guardStart.y == y {
				continue
			}

			guardPosition := guardStart
			currentDirection := North
			currentView := direction[currentDirection]

			testMap := make([]string, len(obsMap))
			copy(testMap, obsMap)
			mark(testMap, Vector{x, y}, '#')

			counter := 0
			for true {
				if guardPosition.x+currentView.x < 0 ||
					guardPosition.x+currentView.x >= len(testMap[0]) ||
					guardPosition.y+currentView.y < 0 ||
					guardPosition.y+currentView.y >= len(testMap) {
					break
				}

				mark(testMap, guardPosition, '^')
				if testMap[guardPosition.y+currentView.y][guardPosition.x+currentView.x] == '#' {
					currentDirection = turnRight(currentDirection)
					currentView = direction[currentDirection]
				} else {
					guardPosition.x += currentView.x
					guardPosition.y += currentView.y
				}
				if counter >= 130*130 {
					possibilities++
					break
				}
				counter++
			}
		}
	}

	fmt.Println("Possibilities: ", possibilities)
}

func predictPath(obsMap []string) []string {
	exploredPlaces := 1
	currentDirection := North
	currentView := direction[currentDirection]

	guardPosition := locateGuard(obsMap)
	if guardPosition.x < 0 || guardPosition.y < 0 {
		panic("Guard not found")
	}

	for true {
		if guardPosition.x+currentView.x < 0 ||
			guardPosition.x+currentView.x >= len(obsMap[0]) ||
			guardPosition.y+currentView.y < 0 ||
			guardPosition.y+currentView.y >= len(obsMap) {
			break
		}

		switch obsMap[guardPosition.y+currentView.y][guardPosition.x+currentView.x] {
		case '#':
			currentDirection = turnRight(currentDirection)
			currentView = direction[currentDirection]
		case 'X':
			guardPosition.x += currentView.x
			guardPosition.y += currentView.y
		default:
			guardPosition.x += currentView.x
			guardPosition.y += currentView.y
			exploredPlaces++
			mark(obsMap, guardPosition, 'X')
		}
	}

	fmt.Println("The Guard explored ", exploredPlaces, " place(s)")
	return obsMap
}

func locateGuard(obsMap []string) Vector {
	for y := range obsMap {
		if strings.Index(obsMap[y], "^") > 0 {
			return Vector{strings.Index(obsMap[y], "^"), y}
		}
	}

	return Vector{-1, -1}
}

func turnRight(direction Direction) Direction {
	switch direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	default:
		return North
	}
}

func mark(obsMap []string, position Vector, char byte) {
	b := []byte(obsMap[position.y])
	b[position.x] = char
	obsMap[position.y] = string(b)
}
