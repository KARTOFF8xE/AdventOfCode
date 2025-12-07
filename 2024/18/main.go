package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

type Node struct {
	coords    Vector
	direction Vector
	steps     int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fallingBytes := []Vector{}
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}

		fallingBytes = append(fallingBytes, Vector{x, y})
	}

	findPath(fallingBytes)
	findPreventingByte(fallingBytes)
}

func findPath(fallingBytes []Vector) {
	size := 71
	field := getField(fallingBytes, size, 1500)

	endPos := Vector{size - 1, size - 1}
	steps := useDijkstra(Vector{0, 0}, endPos, Vector{x: 1, y: 0}, field)
	fmt.Println("Minimal Steps in Maze:", steps)
}

func findPreventingByte(fallingBytes []Vector) {
	size := 71

	for i := 1024; i < len(fallingBytes); i++ {
		field := getField(fallingBytes, size, i)
		endPos := Vector{size - 1, size - 1}
		if useDijkstra(Vector{0, 0}, endPos, Vector{x: 1, y: 0}, field) == 0 {
			fmt.Println("End unreachable after", i, "ns cause of:", fallingBytes[i-1])
			break
		}
	}
}

func getField(fallingBytes []Vector, size, nrOfBytes int) [][]byte {
	field := [][]byte{}

	for _ = range size {
		line := []byte{}
		for _ = range size {
			line = append(line, '.')
		}
		field = append(field, line)
	}

	for _, byte := range fallingBytes[:nrOfBytes] {
		field[byte.y][byte.x] = '#'
	}

	return field
}

func useDijkstra(start, end, direction Vector, fields [][]byte) int {
	todoMap := make(map[Vector]Node)
	todoMap[start] = Node{coords: start, direction: Vector{1, 0}, steps: 0}
	doneMap := make(map[Vector]Node)

	for true {
		if len(todoMap) == 0 {
			return doneMap[end].steps
		}
		var min *Node = nil
		for _, n := range todoMap {
			if min == nil {
				min = &n
			}
			if n.steps < min.steps {
				min = &n
			}
		}
		node := Node{}
		node = *min
		delete(todoMap, node.coords)
		doneMap[node.coords] = node

		watchField(fields, &node, turnLeft(node.direction), end, &todoMap, &doneMap)
		watchField(fields, &node, turnRight(node.direction), end, &todoMap, &doneMap)
		watchField(fields, &node, node.direction, end, &todoMap, &doneMap)
	}
	return doneMap[end].steps
}

func watchField(fields [][]byte, node *Node, direction, end Vector, todoMap, doneMap *map[Vector]Node) {
	watchAt := addVectors(node.coords, direction)
	if watchAt.x < 0 || watchAt.x >= len(fields[0]) ||
		watchAt.y < 0 || watchAt.y >= len(fields[0]) {
		return
	}
	if fields[watchAt.y][watchAt.x] == '.' {
		if _, ok := (*doneMap)[watchAt]; ok {
			if (*doneMap)[watchAt].steps > node.steps+1 {
				(*doneMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + 1}
			}
		} else {
			if _, ok := (*todoMap)[watchAt]; ok {
				if (*todoMap)[watchAt].steps > node.steps+1 {
					(*todoMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + 1}
				}
			} else {
				if watchAt != end {
					(*todoMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + 1}
				} else {
					(*doneMap)[end] = Node{coords: end, direction: direction, steps: node.steps + 1}
				}
			}
		}
	}
}

func addVectors(a, b Vector) Vector {
	return Vector{
		x: a.x + b.x,
		y: a.y + b.y,
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
