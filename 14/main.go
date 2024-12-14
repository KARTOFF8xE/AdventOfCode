package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x int
	y int
}

type Robot struct {
	position Vector
	velocity Vector
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	robots := []Robot{}
	for scanner.Scan() {
		line := scanner.Text()
		robot := Robot{}
		robot.position.x, err = strconv.Atoi(line[len("p="):strings.Index(line, ",")])
		if err != nil {
			panic(err)
		}
		robot.position.y, err = strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, " ")])
		if err != nil {
			panic(err)
		}
		line = line[strings.Index(line, " "):]
		robot.velocity.x, err = strconv.Atoi(line[strings.Index(line, "v=")+len("v=") : strings.Index(line, ",")])
		if err != nil {
			panic(err)
		}
		robot.velocity.y, err = strconv.Atoi(line[strings.Index(line, ",")+1:])
		if err != nil {
			panic(err)
		}

		robots = append(robots, robot)
	}
	defer file.Close()

	simulateRobots(robots)
	searchChristimasTree(robots)
}

func simulateRobots(robots []Robot) {
	quadrants := make([]int, 4)
	playgroundSize := Vector{101, 103}

	for _, robot := range robots {
		for _ = range 100 {
			robot.position.x += robot.velocity.x
			robot.position.y += robot.velocity.y

			if robot.position.x < 0 {
				robot.position.x += playgroundSize.x
			}
			if robot.position.x >= playgroundSize.x {
				robot.position.x -= playgroundSize.x
			}
			if robot.position.y < 0 {
				robot.position.y += playgroundSize.y
			}
			if robot.position.y >= playgroundSize.y {
				robot.position.y -= playgroundSize.y
			}
		}

		if robot.position.x > playgroundSize.x/2 &&
			robot.position.y < playgroundSize.y/2 {
			quadrants[0]++
		}
		if robot.position.x < playgroundSize.x/2 &&
			robot.position.y < playgroundSize.y/2 {
			quadrants[1]++
		}
		if robot.position.x < playgroundSize.x/2 &&
			robot.position.y > playgroundSize.y/2 {
			quadrants[2]++
		}
		if robot.position.x > playgroundSize.x/2 &&
			robot.position.y > playgroundSize.y/2 {
			quadrants[3]++
		}
	}
	safetyFactor := 1
	for i := range quadrants {
		safetyFactor *= quadrants[i]
	}

	fmt.Println("The safety factor is", safetyFactor)
}

func searchChristimasTree(robots []Robot) {
	playgroundSize := Vector{101, 103}

	for i := range 10000 {
		searchForTree(playgroundSize, robots, i)

		for j := range robots {
			robots[j].position.x += robots[j].velocity.x
			robots[j].position.y += robots[j].velocity.y

			if robots[j].position.x < 0 {
				robots[j].position.x += playgroundSize.x
			}
			if robots[j].position.x >= playgroundSize.x {
				robots[j].position.x -= playgroundSize.x
			}
			if robots[j].position.y < 0 {
				robots[j].position.y += playgroundSize.y
			}
			if robots[j].position.y >= playgroundSize.y {
				robots[j].position.y -= playgroundSize.y
			}
		}

	}
}

func searchForTree(playgroundSize Vector, robots []Robot, i int) {
	matrix := make([][]byte, 0)
	for _ = range playgroundSize.y {
		line := []byte{}
		for _ = range playgroundSize.x {
			line = append(line, '.')
		}
		matrix = append(matrix, line)
	}
	for _, robot := range robots {
		matrix[robot.position.y][robot.position.x] = '#'
	}
	filter := [][]byte{
		{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
	}
	if foundPattern(matrix, filter) {
		draw(robots, playgroundSize, i)
		fmt.Println("I found something after", i, "seconds. Look inside imgs/")
	}
}

func foundPattern(matrix, filter [][]byte) bool {
	matrixX := len(matrix[0])
	matrixY := len(matrix)
	filterX := len(filter[0])
	filterY := len(filter)

	match := 0
	for y := range filterY {
		for x := range filterX {
			match += int(filter[y][x]) * int(filter[y][x])
		}
	}

	for y := range matrixY - filterY + 1 {
		for x := range matrixX - filterX + 1 {
			test := 0
			for fy := range filterY {
				for fx := range filterX {
					test += int(matrix[y+fy][x+fx]) * int(filter[fy][fx])
				}
			}

			if test == match {
				return true
			}
		}
	}

	return false
}

func draw(robots []Robot, playgroundSize Vector, index int) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{playgroundSize.x, playgroundSize.y}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	cyan := color.RGBA{100, 200, 200, 0xff}

	for x := range playgroundSize.x {
		for y := range playgroundSize.y {
			img.Set(x, y, color.Black)
		}
	}

	for _, robot := range robots {
		img.Set(robot.position.x, robot.position.y, cyan)
	}

	f, _ := os.Create(fmt.Sprintf("imgs/image_%v.png", index))
	png.Encode(f, img)
}
