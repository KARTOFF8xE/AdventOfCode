package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	compactFilesByFragmentation(lines[0])
	compactFilesByMovingFiles(lines[0])
}

func compactFilesByFragmentation(diskMap string) {
	numbers := extractNumbers(diskMap)

	extDiskMap, filespace := extendDiskMap(numbers)

	for i, number := range extDiskMap {
		if i >= filespace {
			break
		}
		if number == nil {
			extDiskMap[i] = popLastValid(extDiskMap)
		}
	}

	checksum := calcChecksum(extDiskMap)
	fmt.Println("new checksum for fragmentation:", checksum)
}

func compactFilesByMovingFiles(diskMap string) {
	numbers := extractNumbers(diskMap)

	extDiskMap, _ := extendDiskMap(numbers)

	for i := len(numbers) - 1; i >= 0; i-- {
		if !isEven(i) {
			continue
		}

		findFreeSpaceAndMove(i, numbers, extDiskMap)
	}

	checksum := calcChecksum(extDiskMap)
	fmt.Printf("new checksum for moving files: %v\n", checksum)
}

func findFreeSpaceAndMove(i int, numbers []int, extDiskMap []*int) {
	id := 0
	for l := range i {
		id += numbers[l]
	}

	fileSpace := numbers[i]
	dummySpace := make([]*int, fileSpace)

	for j := 0; j < id; j++ {
		if reflect.DeepEqual(extDiskMap[j:j+len(dummySpace)], dummySpace) {
			for k := range fileSpace {
				extDiskMap[j+k] = pointer(i / 2)
			}
			if id > j {
				for l := range fileSpace {
					extDiskMap[id+l] = nil
				}
			}
			return
		}
	}
}

func isEven(number int) bool {
	return number%2 == 0
}

func popLastValid(numbers []*int) *int {
	for i := len(numbers) - 1; i >= 0; i-- {
		if numbers[i] != nil {
			tmp := numbers[i]
			numbers[i] = nil
			return tmp
		}
	}

	return nil
}

func pointer(number int) *int {
	return &number
}

func extractNumbers(diskMap string) []int {
	numbers := []int{}
	for _, digit := range diskMap {
		number := int(digit - '0')
		numbers = append(numbers, number)
	}
	return numbers
}

func extendDiskMap(numbers []int) ([]*int, int) {
	filespace := 0
	extDiskMap := []*int{}
	for i, value := range numbers {
		for _ = range value {
			if isEven(i) {
				extDiskMap = append(extDiskMap, pointer(i/2))
				filespace++
				continue
			}
			extDiskMap = append(extDiskMap, nil)
		}
	}
	return extDiskMap, filespace
}

func calcChecksum(extDiskMap []*int) int64 {
	checksum := int64(0)
	for i, value := range extDiskMap {
		if value != nil {
			checksum += int64(*value * i)
		}
	}
	return checksum
}
