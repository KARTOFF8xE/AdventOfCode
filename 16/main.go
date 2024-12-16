package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

	fields := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scan := scanner.Bytes()
		line := make([]byte, len(scan))
		copy(line, scan)
		fields = append(fields, line)
	}

	search(fields)
}

func localize(fields [][]byte, char byte) Vector {
	for y := range fields {
		if bytes.Contains(fields[y], []byte{char}) {
			return Vector{x: bytes.Index(fields[y], []byte{char}), y: y}
		}
	}
	panic("Reindeer not found")
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

func search(fields [][]byte) {
	endPos := localize(fields, 'E')
	fields[endPos.y][endPos.x] = '.'
	steps := useDijkstra(localize(fields, 'S'), endPos, Vector{x: 1, y: 0}, fields)
	fmt.Println("Minimal Steps in Maze:", steps)
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

		//--------------
		// time.Sleep(time.Millisecond * 5)
		// cmd := exec.Command("clear") //Linux example, its tested
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		// for _, element := range doneMap {
		// 	fields[element.coords.y][element.coords.x] = 'U'
		// }
		// for _, line := range fields {
		// 	fmt.Println(string(line))
		// }
		//--------------

		watchField(fields, &node, turnLeft(node.direction), end, 1001, &todoMap, &doneMap)
		watchField(fields, &node, turnRight(node.direction), end, 1001, &todoMap, &doneMap)
		watchField(fields, &node, node.direction, end, 1, &todoMap, &doneMap)
	}
	return doneMap[end].steps
}

func watchField(fields [][]byte, node *Node, direction, end Vector, inc int, todoMap, doneMap *map[Vector]Node) {
	watchAt := addVectors(node.coords, direction)
	if fields[watchAt.y][watchAt.x] == '.' {
		if _, ok := (*doneMap)[watchAt]; ok {
			if (*doneMap)[watchAt].steps > node.steps+inc {
				(*doneMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + inc}
			}
		} else {
			if _, ok := (*todoMap)[watchAt]; ok {
				if (*todoMap)[watchAt].steps > node.steps+inc {
					(*todoMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + inc}
				}
			} else {
				if watchAt != end {
					(*todoMap)[watchAt] = Node{coords: watchAt, direction: direction, steps: node.steps + inc}
				} else {
					(*doneMap)[end] = Node{coords: end, direction: direction, steps: node.steps + inc}
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
