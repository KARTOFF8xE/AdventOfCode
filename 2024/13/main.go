package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x int64
	y int64
}

type Equation struct {
	BtnA Vector
	BtnB Vector
	Sltn Vector
}

type LinearEq struct {
	m float64
	n float64
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	equations := []Equation{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		eq := Equation{}
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		eq.BtnA = parseBtn(line)

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}
		eq.BtnB = parseBtn(line)

		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		eq.Sltn = parseSltn(line)

		equations = append(equations, eq)
	}

	play(equations)
	playWithALotMoney(equations)
}

func parseBtn(line string) Vector {
	x, err := strconv.ParseInt(line[len("Button X: X+"):strings.Index(line, ",")], 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(line[strings.Index(line, "Y+")+len("Y+"):], 10, 64)
	if err != nil {
		panic(err)
	}

	return Vector{x: x, y: y}
}

func parseSltn(line string) Vector {
	x, err := strconv.ParseInt(line[len("Prize: X="):strings.Index(line, ",")], 10, 64)
	if err != nil {
		panic(err)
	}
	y, err := strconv.ParseInt(line[strings.Index(line, "Y=")+len("Y="):], 10, 64)
	if err != nil {
		panic(err)
	}

	return Vector{x: x, y: y}
}

func play(equations []Equation) {
	token := 0
	rayA := []Vector{}
	rayB := []Vector{}
	for _, equation := range equations {
		rayA = getRay(Vector{0, 0}, equation.BtnA, 1)
		rayB = getRay(equation.Sltn, equation.BtnB, -1)
		nrBtnAPressed, nrBtnBPressed := getIntersectionStampIndex(rayA, rayB)
		token += 3*nrBtnAPressed + nrBtnBPressed
	}

	fmt.Println("used", token, "tokens to win all possible prices")
}

func getRay(startVec, dirVec Vector, direction int) []Vector {
	ray := []Vector{}

	for i := range 101 {
		ray = append(ray, Vector{
			x: startVec.x + int64(direction*i)*dirVec.x,
			y: startVec.y + int64(direction*i)*dirVec.y,
		})
	}
	return ray
}

func getIntersectionStampIndex(rayA, rayB []Vector) (int, int) {
	for i := range rayA {
		for j := range rayB {
			if rayA[i] == rayB[j] {
				return i, j
			}
		}
	}

	return 0, 0
}

func playWithALotMoney(equations []Equation) {
	token := int64(0)
	for _, equation := range equations {
		equation.Sltn.x += 10000000000000
		equation.Sltn.y += 10000000000000
		x, y := getIntersection(equation)
		nrBtnAPressed, nrBtnBPressed := analyseIntersection(equation, x, y)
		token += 3*nrBtnAPressed + nrBtnBPressed

	}
	fmt.Println("used", token, "tokens to win all possible prices")
}

func getIntersection(equation Equation) (float64, float64) {
	a := LinearEq{
		m: float64(equation.BtnA.y) / float64(equation.BtnA.x),
		n: 0,
	}
	b := LinearEq{
		m: float64(equation.BtnB.y) / float64(equation.BtnB.x),
		n: 0,
	}
	b.n = float64(equation.Sltn.y) - b.m*float64(equation.Sltn.x)

	x := b.n / (a.m - b.m)
	y := a.m * x

	return roundFloatFault(x), roundFloatFault(y)
}

func analyseIntersection(equation Equation, x, y float64) (int64, int64) {
	if isIntegral(roundFloatFault(x/float64(equation.BtnA.x))) &&
		isIntegral(roundFloatFault(y/float64(equation.BtnA.y))) &&
		isIntegral(roundFloatFault((float64(equation.Sltn.x)-x)/float64(equation.BtnB.x))) &&
		isIntegral(roundFloatFault((float64(equation.Sltn.y)-y)/float64(equation.BtnB.y))) {
		foo := x / float64(equation.BtnA.x)
		bar := (float64(equation.Sltn.y) - y) / float64(equation.BtnB.y)

		return int64(foo), int64(bar)
	}
	return 0, 0
}

func isIntegral(val float64) bool {
	return val == float64(int64(val))
}

func roundFloatFault(num float64) float64 {
	str := strconv.FormatFloat(num, 'f', -1, 64)

	parts := strings.Split(str, ".")
	if len(parts) < 2 {
		return num
	}

	decimalPlaces := len(parts[1])

	factor := math.Pow(10, float64(decimalPlaces-3))
	return math.Round(num*factor) / factor
}
