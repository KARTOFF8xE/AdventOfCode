package main

import (
	"bufio"
	"fmt"
	"os"
)

type direction = int

type Vector struct {
	x int
	y int
}

type Trailhead struct {
	coords Vector
	summit map[int]map[int]int
}

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	topMap := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		intLine := []int{}
		for _, char := range line {
			intLine = append(intLine, int(char)-'0')
		}
		topMap = append(topMap, intLine)
	}

	findHikes(topMap)
}

func findHikes(topMap [][]int) {
	trailheads := getTrailheads(topMap)

	for i := range trailheads {
		lookForPath(topMap, trailheads[i].coords, &trailheads[i])
	}

	generalScore := 0
	generalRaiting := 0
	for _, trailhead := range trailheads {
		for x, ys := range trailhead.summit {
			if trailhead.summit[x] == nil {
				continue
			}
			generalScore += len(trailhead.summit[x])
			for _, y := range ys {
				generalRaiting += y
			}
		}
	}
	fmt.Println("Found", generalScore, " trails")
	fmt.Println("Found", generalRaiting, " distinct trails")
}

func lookForPath(topMap [][]int, currentPos Vector, trailhead *Trailhead) {
	if topMap[currentPos.y][currentPos.x] == 9 {
		if trailhead.summit[currentPos.x] == nil {
			trailhead.summit[currentPos.x] = make(map[int]int)
		}
		trailhead.summit[currentPos.x][currentPos.y]++
		return
	}

	for i := 0; i < 4; i++ {
		direction := getVectorByDirection(i)
		if currentPos.x+direction.x < len(topMap[0]) &&
			currentPos.x+direction.x >= 0 &&
			currentPos.y+direction.y < len(topMap) &&
			currentPos.y+direction.y >= 0 {
			if topMap[currentPos.y][currentPos.x]+1 == topMap[currentPos.y+direction.y][currentPos.x+direction.x] {
				lookForPath(
					topMap,
					Vector{
						x: currentPos.x + direction.x,
						y: currentPos.y + direction.y,
					},
					trailhead,
				)
			}
		}
	}
}

func getTrailheads(topMap [][]int) []Trailhead {
	trailheads := make([]Trailhead, 0)

	for y := range topMap {
		for x := range topMap[y] {
			if topMap[y][x] == 0 {
				trailheads = append(trailheads, Trailhead{coords: Vector{x: x, y: y}, summit: make(map[int]map[int]int)})
			}
		}
	}

	return trailheads
}

func getVectorByDirection(direction direction) Vector {
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
		panic("unkown direction")
	}
}
