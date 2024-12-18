package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	regA int64
	regB int64
	regC int64

	progPtr int

	program []int
	output  []int64
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	computer := Computer{progPtr: 0}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	parseReg(scanner, &computer.regA)
	parseReg(scanner, &computer.regB)
	parseReg(scanner, &computer.regC)

	scanner.Scan()
	line := scanner.Text()
	programmString := strings.Split(line, " ")[1]
	programmStrings := strings.Split(programmString, ",")
	for _, item := range programmStrings {
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}

		computer.program = append(computer.program, itemInt)
	}

	computer.determineOutput()
}

func parseReg(scanner *bufio.Scanner, reg *int64) {
	line := scanner.Text()
	regString := strings.Split(line, " ")
	regInt, err := strconv.ParseInt(regString[2], 10, 64)
	*reg = regInt
	if err != nil {
		panic(err)
	}
	scanner.Scan()
}

func (c *Computer) determineOutput() {
	for ok := true; ok; ok = c.progPtr < len(c.program) {
		switch c.program[c.progPtr] {
		case 0:
			c.adv(c.program[c.progPtr+1])
			c.progPtr += 2
		case 1:
			c.bxl(c.program[c.progPtr+1])
			c.progPtr += 2
		case 2:
			c.bst(c.program[c.progPtr+1])
			c.progPtr += 2
		case 3:
			c.jnz(c.program[c.progPtr+1])
		case 4:
			c.bxc(c.program[c.progPtr+1])
			c.progPtr += 2
		case 5:
			c.out(c.program[c.progPtr+1])
			c.progPtr += 2
		case 6:
			c.bdv(c.program[c.progPtr+1])
			c.progPtr += 2
		case 7:
			c.cdv(c.program[c.progPtr+1])
			c.progPtr += 2
		default:
		}
	}
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(c.output)), ","), "[]"))
}

func (c *Computer) adv(operand int) {
	op := int64(operand)
	c.evalOperand(&op)

	c.regA = c.regA / int64(math.Pow(2, float64(op)))
}

func (c *Computer) bxl(operand int) {
	c.regB = c.regB ^ int64(operand)
}

func (c *Computer) bst(operand int) {
	op := int64(operand)
	c.evalOperand(&op)

	c.regB = 7 & (op % 8)
}

func (c *Computer) jnz(operand int) {
	if c.regA == 0 {
		c.progPtr += 2
		return
	}

	c.progPtr = operand
}

func (c *Computer) bxc(operand int) {
	c.regB = c.regB ^ c.regC
}

func (c *Computer) out(operand int) {
	op := int64(operand)
	c.evalOperand(&op)

	c.output = append(c.output, op%8)
}

func (c *Computer) bdv(operand int) {
	op := int64(operand)
	c.evalOperand(&op)

	c.regB = c.regA / int64(math.Pow(2, float64(op)))
}

func (c *Computer) cdv(operand int) {
	op := int64(operand)
	c.evalOperand(&op)

	c.regC = c.regA / int64(math.Pow(2, float64(op)))
}

func (c *Computer) evalOperand(operand *int64) {
	switch *operand {
	case 4:
		*operand = c.regA
	case 5:
		*operand = c.regB
	case 6:
		*operand = c.regC
	case 7:
		panic("forbitten State")
	default:
	}
}
