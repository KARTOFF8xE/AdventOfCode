package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	edgesFrom []int
	edgesTo   []int
	steps     *int
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scan := scanner.Text()
	towels := strings.Split(scan, ", ")

	designs := []string{}
	for scanner.Scan() {
		scan := scanner.Text()
		if len(scan) == 0 {
			continue
		}
		designs = append(designs, scan)
	}

	possibleDesigns(towels, designs)
	sumPossibleDesignDesigns(towels, designs)
}

func sumPossibleDesignDesigns(towels, designs []string) {
	counter := int64(0)
	for _, design := range designs {
		memo := make(map[string]int64)
		counter += search(towels, design, memo)
	}
	fmt.Printf("Found %v ways to design all designs\n", counter)
}

func search(towels []string, design string, memo map[string]int64) int64 {
	if val, ok := memo[design]; ok {
		return val
	}
	if len(design) == 0 {
		return 1
	}

	counter := int64(0)
	for _, towel := range towels {
		if len(towel) > len(design) {
			continue
		}
		if strings.Index(design, towel) == 0 {
			counter += search(towels, design[len(towel):], memo)
		}
	}

	memo[design] = counter

	return counter
}

func possibleDesigns(towels, designs []string) {
	counterPosDes := 0

	for _, design := range designs {
		indexMap := createIndexMap(design, towels)
		graph := make(map[int]Node)
		for _, elements := range indexMap {
			for _, element := range elements {
				edgesTo := graph[element[0]].edgesTo
				edgesTo = append(edgesTo, element[1])
				graph[element[0]] = Node{edgesTo: edgesTo, steps: nil}
			}
		}

		if djikstra(graph, 0, len(design)-1) != nil {
			counterPosDes++
		}
	}

	fmt.Println("There are", counterPosDes, "designs possible")
}

func createIndexMap(design string, towels []string) map[string][][]int {
	indexMap := make(map[string][][]int)

	for _, towel := range towels {
		match, err := regexp.Compile(string(towel))
		if err != nil {
			panic(err)
		}
		indexes := match.FindAllStringIndex(design, -1)
		if len(indexes) != 0 {
			indexMap[towel] = indexes
		}
	}

	return indexMap
}

func djikstra(graph map[int]Node, start, end int) *int {
	todoMap := make(map[int]Node)
	doneMap := make(map[int]Node)

	startNode := graph[0]
	startNode.steps = POINTER(0)
	todoMap[0] = startNode

	for true {
		if len(todoMap) == 0 {
			return nil
		}

		key := -1
		var min *Node = nil
		for k, n := range todoMap {
			if min == nil {
				key = k
				min = &n
			}

			if n.steps != nil && *n.steps > *min.steps {
				key = k
				min = &n
			}
		}

		node := *min
		doneMap[key] = *min
		delete(todoMap, key)

		reachedEnd := evolveNode(graph, node, end, &todoMap, &doneMap)
		if reachedEnd {
			break
		}
	}
	return doneMap[end].steps
}

func evolveNode(graph map[int]Node, node Node, end int, todoMap, doneMap *map[int]Node) bool {
	for _, element := range node.edgesTo {
		if _, ok := (*doneMap)[element]; ok {
			if *node.steps+1 < *(*doneMap)[element].steps {
				(*doneMap)[element] = Node{edgesTo: (*doneMap)[element].edgesTo, steps: POINTER(*node.steps + 1)}
			}
		} else {
			if _, ok := (*todoMap)[element]; ok {
				if *node.steps+1 < *(*todoMap)[element].steps {
					(*todoMap)[element] = Node{edgesTo: (*todoMap)[element].edgesTo, steps: POINTER(*node.steps + 1)}
				}
			} else {
				if element == end {
					if node.steps == nil {
						fmt.Println("STEPS IS NILL")
					} else {
						(*doneMap)[element] = Node{edgesTo: []int{}, steps: POINTER(*node.steps + 1)}
					}
					return true
				} else {
					n := graph[element]
					n.steps = POINTER(*node.steps + 1)
					(*todoMap)[element] = n
				}
			}
		}
	}
	return false
}

func POINTER[T any](val T) *T {
	return &val
}
