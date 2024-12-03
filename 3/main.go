package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := getInput()

	getSumOfProducts(input)
	getPreciseSumOfProducts(input)
}

func getSumOfProducts(input string) {
	regex, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	if err != nil {
		fmt.Println("invalid regex with Error:", err)
		return
	}
	validOps := regex.FindAllString(input, -1)

	regex, err = regexp.Compile("\\d{1,3},\\d{1,3}")
	if err != nil {
		fmt.Println("invalid regex with Error:", err)
		return
	}

	sumOfProducts := 0
	for _, item := range validOps {
		factors := strings.Split(regex.FindString(item), ",")
		factor1, err1 := strconv.Atoi(factors[0])
		factor2, err2 := strconv.Atoi(factors[1])
		if err1 != nil || err2 != nil {
			fmt.Println("at least one invalid factor")
			return
		}
		sumOfProducts += factor1 * factor2
	}

	fmt.Println("sumOfProducts:", sumOfProducts)
}

func getPreciseSumOfProducts(input string) {
	regexDo, err := regexp.Compile("do\\(\\)")
	if err != nil {
		fmt.Println("invalid regex with Error:", err)
		return
	}
	regexDont, err := regexp.Compile("don't\\(\\)")
	if err != nil {
		fmt.Println("invalid regex with Error:", err)
		return
	}

	indexDo := 0
	indexDont := 0
	valids := ""

	for indexDo != len(input) {
		i := regexDont.FindStringIndex(input)
		if len(i) != 0 {
			indexDont = i[1] - 1
		} else {
			indexDont = len(input)
		}

		if indexDo < indexDont {
			valids = fmt.Sprint(valids, input[indexDo:indexDont])
		}
		input = input[indexDont:]

		i = regexDo.FindStringIndex(input)
		if len(i) != 0 {
			indexDo = i[0]
		} else {
			indexDo = len(input)
		}
	}

	fmt.Print("only valid mul()'s: ")
	getSumOfProducts(valids)
}
