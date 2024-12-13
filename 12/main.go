package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Direction int

const (
	WEST Direction = iota
	NORTH
	EAST
	SOUTH
)

type Vector struct {
	x int
	y int
}

type RegInfos struct {
	area        int
	borderNorth map[int][]int // y -> []x
	borderEast  map[int][]int // x -> []y
	borderSouth map[int][]int // y -> []x
	borderWest  map[int][]int // x -> []y
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fields := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scan := scanner.Bytes()
		line := make([]byte, len(scan))
		copy(line, scan)
		fields = append(fields, line)
	}

	calcPrice(fields)
}

func calcPrice(fields [][]byte) {
	regCounter := 0
	costStd := 0
	costDisc := 0
	for y := range fields {
		for x := range fields[y] {
			if fields[y][x] < 'A' || fields[y][x] > 'Z' {
				continue
			}
			regInfos := RegInfos{
				area:        0,
				borderNorth: make(map[int][]int),
				borderEast:  make(map[int][]int),
				borderSouth: make(map[int][]int),
				borderWest:  make(map[int][]int),
			}
			surroundRegion(fields, fields[y][x], Vector{x, y}, nil, &regInfos)
			std, disc := analyzeRegion(regInfos)
			costStd += std
			costDisc += disc
			regCounter++
		}
	}
	fmt.Println("Found", regCounter, "different regions")
	fmt.Println("Costs (standard):", costStd)
	fmt.Println("Costs (bulk discount):", costDisc)
}

func surroundRegion(fields [][]byte, regionType byte, currentField Vector, dir *Direction, regInfos *RegInfos) {
	if dir != nil {
		currentField = addVectors(currentField, getDirectionVector(*dir))
	}

	if currentField.x < 0 || currentField.x >= len(fields[0]) ||
		currentField.y < 0 || currentField.y >= len(fields) {
		appendBorders(dir, regInfos, currentField)
		return
	}

	if fields[currentField.y][currentField.x] != regionType {
		if fields[currentField.y][currentField.x] != regionType+32 {
			appendBorders(dir, regInfos, currentField)
		}
		return
	}

	fields[currentField.y][currentField.x] = fields[currentField.y][currentField.x] + 32
	regInfos.area++
	for direction := range Direction(4) {
		surroundRegion(fields, regionType, currentField, &direction, regInfos)
	}
}

func appendBorders(dir *Direction, regInfos *RegInfos, field Vector) {
	switch *dir {
	case NORTH:
		regInfos.borderNorth[field.y] = append(regInfos.borderNorth[field.y], field.x)
	case EAST:
		regInfos.borderEast[field.x] = append(regInfos.borderEast[field.x], field.y)
	case SOUTH:
		regInfos.borderSouth[field.y] = append(regInfos.borderSouth[field.y], field.x)
	case WEST:
		regInfos.borderWest[field.x] = append(regInfos.borderWest[field.x], field.y)
	default:
		panic("unknown direction")
	}
}

func analyzeRegion(regInfos RegInfos) (int, int) {
	perimeterStd := 0
	perimeterDisc := 0
	borders := []map[int][]int{
		regInfos.borderNorth,
		regInfos.borderEast,
		regInfos.borderSouth,
		regInfos.borderWest,
	}
	for _, border := range borders {
		disc, std := analyzePerimeters(border)
		perimeterDisc += disc
		perimeterStd += std
	}

	return perimeterDisc * regInfos.area, perimeterStd * regInfos.area
}

func analyzePerimeters(border map[int][]int) (int, int) {
	perimeterStd := 0
	perimeterDisc := 0
	for _, x := range border {
		perimeterDisc++
		if !slices.IsSorted(x) {
			slices.Sort(x)
		}
		for i := range len(x) - 1 {
			if x[i+1]-x[i] != 1 {
				perimeterDisc++
			}
		}
	}
	for _, v := range border {
		perimeterStd += len(v)
	}

	return perimeterStd, perimeterDisc
}

func addVectors(v1, v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func getDirectionVector(direction Direction) Vector {
	switch direction {
	case NORTH:
		return Vector{0, -1}
	case EAST:
		return Vector{1, 0}
	case SOUTH:
		return Vector{0, 1}
	case WEST:
		return Vector{-1, 0}
	default:
		panic("Unknown direction")
	}
}
