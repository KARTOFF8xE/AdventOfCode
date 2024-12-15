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

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fields := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scan := scanner.Bytes()
		if string(scan) == "" {
			break
		}
		line := make([]byte, len(scan))
		copy(line, scan)
		fields = append(fields, line)
		copy(fields[len(fields)-1], line)
	}

	movements := []byte{}
	for scanner.Scan() {
		line := scanner.Bytes()
		movements = append(movements, line...)
	}

	wideFields := stretchWareHouse(fields)
	predictWareHousePlan(fields, movements)
	predictWideHousePlan(wideFields, movements)
}

func stretchWareHouse(fields [][]byte) [][]byte {
	wideHouse := make([][]byte, 0)
	for y := range fields {
		line := []byte{}
		line = append(line, fields[y]...)
		line = append(line, fields[y]...)
		wideHouse = append(wideHouse, line)
		for x := range fields[y] {
			switch fields[y][x] {
			case '#':
				wideHouse[y][2*x] = '#'
				wideHouse[y][2*x+1] = '#'
			case 'O':
				wideHouse[y][2*x] = '['
				wideHouse[y][2*x+1] = ']'
			case '@':
				wideHouse[y][2*x] = '@'
				wideHouse[y][2*x+1] = '.'
			default:
				wideHouse[y][2*x] = '.'
				wideHouse[y][2*x+1] = '.'
			}
		}
	}

	return wideHouse
}

func predictWareHousePlan(fields [][]byte, movements []byte) {
	robotPos := localizeRobot(fields)
	fields[robotPos.y][robotPos.x] = '.'

	for _, movement := range movements {
		direction := getDirectionfieldByMovement(movement)
		viewField := addVectors(robotPos, direction)

		switch fields[viewField.y][viewField.x] {
		case '.':
			robotPos = viewField
		case 'O':
			vec := isMovable(fields, robotPos, direction)
			if vec != nil {
				fields[vec.y][vec.x] = 'O'
				fields[addVectors(robotPos, direction).y][addVectors(robotPos, direction).x] = '.'
				robotPos = addVectors(robotPos, direction)
			}
		default:
		}
	}
	calcSumOfGPSBoxes(fields, 'O')
}

func localizeRobot(fields [][]byte) Vector {
	for y := range fields {
		if bytes.Contains(fields[y], []byte{'@'}) {
			return Vector{x: bytes.Index(fields[y], []byte{'@'}), y: y}
		}
	}
	panic("Robot not found")
}

func addVectors(a, b Vector) Vector {
	return Vector{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func getDirectionfieldByMovement(movement byte) Vector {
	switch movement {
	case '^':
		return Vector{x: 0, y: -1}
	case '>':
		return Vector{x: 1, y: 0}
	case 'v':
		return Vector{x: 0, y: 1}
	case '<':
		return Vector{x: -1, y: 0}
	default:
		panic("unknown movement")
	}
}

func isMovable(fields [][]byte, robotPos, direction Vector) *Vector {
	robotPos = addVectors(robotPos, direction)
	switch fields[robotPos.y][robotPos.x] {
	case '.':
		return &robotPos
	case 'O':
		return isMovable(fields, robotPos, direction)
	default:
		return nil
	}
}

func calcSumOfGPSBoxes(fields [][]byte, char byte) {
	sum := int64(0)

	for y := range fields {
		for x := range fields[y] {
			if fields[y][x] == char {
				sum += int64(y*100) + int64(x)
			}
		}
	}

	fmt.Println("Sum of GPS coordinates:", sum)
}

func predictWideHousePlan(fields [][]byte, movements []byte) {
	robotPos := localizeRobot(fields)
	fields[robotPos.y][robotPos.x] = '.'

	for _, movement := range movements {
		direction := getDirectionfieldByMovement(movement)
		viewField := addVectors(robotPos, direction)

		switch fields[viewField.y][viewField.x] {
		case '.':
			robotPos = viewField
		case '[':
			fallthrough
		case ']':
			if math.Abs(float64(direction.x)) == 1 { // horizontal
				vec := isHorizontalMovable(fields, robotPos, direction)
				if vec != nil && direction.x == -1 { // to the left
					y := vec.y
					for x := vec.x; x < robotPos.x; x++ {
						fields[y][x] = fields[y][x+1]
					}
				}
				if vec != nil && direction.x == 1 { // to the right
					y := vec.y
					for x := vec.x; x > robotPos.x; x-- {
						fields[y][x] = fields[y][x-1]
					}
				}
				if vec != nil {
					robotPos = addVectors(robotPos, direction)
				}
			} else { // vertical
				if isVerticalMovable(fields, robotPos, direction) {
					moveVertically(fields, robotPos, direction)
					fields[robotPos.y][robotPos.x] = '.'
					fields[robotPos.y][robotPos.x+direction.x] = '.'
					robotPos = addVectors(robotPos, direction)
				}
			}
		default:
		}
	}
	calcSumOfGPSBoxes(fields, '[')
}
func isHorizontalMovable(fields [][]byte, robotPos, direction Vector) *Vector {
	viewPos := addVectors(robotPos, direction)
	switch fields[viewPos.y][viewPos.x] {
	case '.':
		return &viewPos
	case '[':
		fallthrough
	case ']':
		return isHorizontalMovable(fields, viewPos, direction)
	default:
		return nil
	}
}

func isVerticalMovable(fields [][]byte, robotPos, direction Vector) bool {
	viewPos := addVectors(robotPos, direction)
	switch fields[viewPos.y][viewPos.x] {
	case '.':
		return true
	case '[':
		return isVerticalMovable(fields, viewPos, direction) &&
			isVerticalMovable(fields, addVectors(viewPos, getDirectionfieldByMovement('>')), direction)
	case ']':
		return isVerticalMovable(fields, addVectors(viewPos, getDirectionfieldByMovement('<')), direction) &&
			isVerticalMovable(fields, viewPos, direction)
	default:
		return false
	}
}

func moveVertically(fields [][]byte, robotPos, direction Vector) {
	viewPos := addVectors(robotPos, direction)

	if fields[viewPos.y][viewPos.x] == '[' {
		moveVertically(fields, viewPos, direction)
		neighbor := addVectors(viewPos, getDirectionfieldByMovement('>'))
		moveVertically(fields, neighbor, direction)
	}

	if fields[viewPos.y][viewPos.x] == ']' {
		neighbor := addVectors(viewPos, getDirectionfieldByMovement('<'))
		moveVertically(fields, neighbor, direction)
		moveVertically(fields, viewPos, direction)
	}

	if fields[viewPos.y][viewPos.x] == '.' {
		fields[viewPos.y][viewPos.x] = fields[viewPos.y-direction.y][viewPos.x]
		fields[viewPos.y-direction.y][viewPos.x] = '.'
	}
}
