package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Vector struct {
	x int
	y int
}

type Field struct {
	direction Vector
	time      int
}

type Racer struct {
	pos  Vector
	dir  Vector
	time int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fields := [][]byte{}
	for scanner.Scan() {
		scan := scanner.Bytes()
		line := make([]byte, len(scan))
		copy(line, scan)

		fields = append(fields, line)
	}

	raceIt(fields, 2, 100)
	raceIt(fields, 20, 100)
}

func raceIt(fields [][]byte, cheatTime, n int) {
	start := localize(fields, 'S')
	end := localize(fields, 'E')

	direction := Vector{0, -1}

	raceMap := initialWalk(fields, start, direction)

	me := Racer{pos: start, dir: raceMap[start].direction, time: 0}

	counter := 0
	relCheatSpots := getPossibleCheatSpots(cheatTime)

	for me.pos != end {
		fmt.Printf("\r%4.2v%% ", float64(me.time)/float64(raceMap[end].time)*100)

		counter += me.getCheatSpot(raceMap, relCheatSpots, n)

		me.race(raceMap)
	}

	fmt.Printf("%v cheats save more than %vps with a cheattime of %vps\n", counter, n, cheatTime)
}

func localize(fields [][]byte, char byte) Vector {
	for y := range fields {
		if bytes.Contains(fields[y], []byte{char}) {
			return Vector{x: bytes.Index(fields[y], []byte{char}), y: y}
		}
	}
	panic("Reindeer not found")
}

func initialWalk(fields [][]byte, start, direction Vector) map[Vector]Field {
	raceMap := make(map[Vector]Field)

	counter := 0
	walk(fields, start, direction, &raceMap, counter)
	walk(fields, start, turnLeft(direction), &raceMap, counter)
	walk(fields, start, turnRight(direction), &raceMap, counter)

	return raceMap
}

func walk(fields [][]byte, position, direction Vector, raceMap *map[Vector]Field, counter int) {
	watchAt := addVectors(position, direction)

	if fields[watchAt.y][watchAt.x] == 'E' {
		(*raceMap)[position] = Field{direction: direction, time: counter}
		(*raceMap)[watchAt] = Field{direction: Vector{x: 0, y: 0}, time: counter + 1}

		return
	}

	if fields[watchAt.y][watchAt.x] == '.' {
		(*raceMap)[position] = Field{direction: direction, time: counter}
		walk(fields, watchAt, direction, raceMap, counter+1)
		walk(fields, watchAt, turnLeft(direction), raceMap, counter+1)
		walk(fields, watchAt, turnRight(direction), raceMap, counter+1)
	}
}

func getPossibleCheatSpots(ps int) map[Vector]int {
	posCheatSpots := make(map[Vector]int)

	search(Vector{0, 0}, Vector{0, -1}, ps, &posCheatSpots)

	return posCheatSpots
}

func search(pos, direction Vector, ps int, posCheatSpots *map[Vector]int) {
	watchAt := addVectors(pos, direction)
	manhDist := getManhattenDist(Vector{0, 0}, watchAt)
	if !(manhDist <= ps) {
		return
	}

	if _, ok := (*posCheatSpots)[watchAt]; ok {
		return
	}
	(*posCheatSpots)[watchAt] = manhDist

	for _, direction := range []Vector{Vector{0, -1}, Vector{0, 1}, Vector{-1, 0}, Vector{1, 0}} {
		search(watchAt, direction, ps, posCheatSpots)
	}
}

func (r Racer) getCheatSpot(raceMap map[Vector]Field, possibleCheatSpots map[Vector]int, n int) int {
	counter := 0
	for relSpot, _ := range possibleCheatSpots {
		spot := addVectors(r.pos, relSpot)
		if _, ok := raceMap[spot]; ok {
			newTime := r.time + possibleCheatSpots[relSpot]
			if newTime < raceMap[spot].time && raceMap[spot].time-newTime >= 100 {
				counter++
			}
		}
	}

	return counter
}

func getManhattenDist(vec1, vec2 Vector) int {
	return int(math.Abs(float64(vec1.x-vec2.x)) + math.Abs(float64(vec1.y-vec2.y)))
}

func (r *Racer) race(raceMap map[Vector]Field) {
	oldpos := r.pos
	r.pos = addVectors(r.pos, raceMap[r.pos].direction)
	if r.pos != oldpos {
		r.time = r.time + 1
	}
}

func turnRight(direction Vector) Vector {
	switch direction {
	case Vector{x: -1, y: 0}:
		return Vector{x: 0, y: -1}
	case Vector{x: 0, y: -1}:
		return Vector{x: 1, y: 0}
	case Vector{x: 1, y: 0}:
		return Vector{x: 0, y: 1}
	case Vector{x: 0, y: 1}:
		return Vector{x: -1, y: 0}
	default:
		panic("unknown movement")
	}
}

func turnLeft(direction Vector) Vector {
	direction.x = -direction.x
	direction.y = -direction.y
	return turnRight(direction)
}

func addVectors(a, b Vector) Vector {
	return Vector{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}
