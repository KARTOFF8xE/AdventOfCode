package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
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

	codes := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code := scanner.Text()
		codes = append(codes, code)
	}

	controllRobot(codes)
}

func controllRobot(codes []string) {
	numPadFields := getNumPad()
	numPad := make(map[byte]Vector)
	for _, char := range []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A'} {
		numPad[char] = localize(numPadFields, char)
	}

	keyPadFields := getKeyPad()
	keyPad := make(map[byte]Vector)
	for _, char := range []byte{'^', '>', 'v', '<', 'A'} {
		keyPad[char] = localize(keyPadFields, char)
	}

	maxDepth := 5
	proceedure := make([][]byte, maxDepth)
	state := slices.Repeat([]byte{'A'}, maxDepth)
	for _, code := range codes {
		for _, element := range code {
			useNumPad(numPad[state[0]], numPad[byte(element)], keyPad, maxDepth, state, proceedure)
			state[0] = byte(element)
		}
	}

	for i, line := range proceedure {
		fmt.Printf("%v: %v\n", i, string(line))
	}
}

func useNumPad(current, dst Vector, keyPad map[byte]Vector, maxDepth int, state []byte, proceedure [][]byte) {
	moveVertical := dst.y - current.y
	for i := moveVertical; i != 0; i -= moveVertical / int(math.Abs(float64(moveVertical))) {
		if moveVertical < 0 {
			useKeyPad('^', keyPad, 1, maxDepth, state, proceedure)
			proceedure[0] = append(proceedure[0], '^')
		}
		if moveVertical > 0 {
			useKeyPad('v', keyPad, 1, maxDepth, state, proceedure)
			proceedure[0] = append(proceedure[0], 'v')
		}
	}

	moveHorizontal := dst.x - current.x
	for i := moveHorizontal; i != 0; i -= moveHorizontal / int(math.Abs(float64(moveHorizontal))) {
		if moveHorizontal < 0 {
			useKeyPad('<', keyPad, 1, maxDepth, state, proceedure)
			proceedure[0] = append(proceedure[0], '<')
		}
		if moveHorizontal > 0 {
			useKeyPad('>', keyPad, 1, maxDepth, state, proceedure)
			proceedure[0] = append(proceedure[0], '>')
		}
	}

	proceedure[0] = append(proceedure[0], 'A')
}

func useKeyPad(dst byte, keyPad map[byte]Vector, depth, maxDepth int, state []byte, proceedure [][]byte) {
	if depth == maxDepth {
		return
	}

	moveVertical := keyPad[dst].y - keyPad[state[depth]].y
	for i := moveVertical; i != 0; i -= moveVertical / int(math.Abs(float64(moveVertical))) {
		if moveVertical < 0 {
			useKeyPad('^', keyPad, depth+1, maxDepth, state, proceedure)
			proceedure[depth] = append(proceedure[depth], '^')
		}
		if moveVertical > 0 {
			useKeyPad('v', keyPad, depth+1, maxDepth, state, proceedure)
			proceedure[depth] = append(proceedure[depth], 'v')
		}
	}

	moveHorizontal := keyPad[dst].x - keyPad[state[depth]].x
	for i := moveHorizontal; i != 0; i -= moveHorizontal / int(math.Abs(float64(moveHorizontal))) {
		if moveHorizontal < 0 {
			useKeyPad('<', keyPad, depth+1, maxDepth, state, proceedure)
			proceedure[depth] = append(proceedure[depth], '<')
		}
		if moveHorizontal > 0 {
			useKeyPad('>', keyPad, depth+1, maxDepth, state, proceedure)
			proceedure[depth] = append(proceedure[depth], '>')
		}
	}

	proceedure[depth] = append(proceedure[depth], 'A')
	state[depth] = dst
}

func getNumPad() []string {
	return []string{
		"789",
		"456",
		"123",
		".0A",
	}
}

func getKeyPad() []string {
	return []string{
		".^A",
		"<v>",
	}
}

func localize(fields []string, char byte) Vector {
	for y, line := range fields {
		x := strings.Index(line, string(char))
		if x != -1 {
			return Vector{x: x, y: y}
		}
	}

	panic("invalid char")
}
