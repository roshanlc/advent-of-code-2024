package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const day = "day-5"
const fileName = day + "-input.txt"

type Pair struct {
	First  int
	Second int
}

func main() {
	// open file
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`(?m)^$`)
	splits := re.Split(string(file), -1)

	fPairs, err := extractFirstPairs(splits[0])
	if err != nil {
		log.Fatal(err)
	}

	sPairs, err := extractSecondPairs(splits[1])
	if err != nil {
		log.Fatal(err)
	}

	// custom sorting function based on first input section
	cmp := func(a, b int) int {
		for _, p := range fPairs {
			if p.First == a && p.Second == b {
				return -1
			}
		}
		return 0
	}

	// loop through second input section
	run := func(sorted bool) int {
		midTotal := 0
		for _, s := range sPairs {
			if slices.IsSortedFunc(s, cmp) == sorted {
				slices.SortFunc(s, cmp)
				midTotal += s[len(s)/2]
			}
		}

		return midTotal
	}

	part1 := run(true)
	part2 := run(false)

	fmt.Println("[Part 1] Total:", part1)
	fmt.Println("[Part 2] Total:", part2)
}

func extractFirstPairs(fileContents string) ([]Pair, error) {
	var pairs []Pair

	scanner := bufio.NewScanner(strings.NewReader(fileContents))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, "|")
		if len(splits) != 2 {
			continue
		}

		f, _ := strconv.Atoi(splits[0])
		s, _ := strconv.Atoi(splits[1])
		pairs = append(pairs, Pair{First: f, Second: s})
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return pairs, nil
}

func extractSecondPairs(fileContents string) ([][]int, error) {
	var pairs [][]int

	scanner := bufio.NewScanner(strings.NewReader(fileContents))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		splits := strings.Split(text, ",")
		if len(splits) < 3 {
			continue
		}

		var p []int
		for _, n := range splits {
			num, _ := strconv.Atoi(n)
			p = append(p, num)
		}

		pairs = append(pairs, p)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return pairs, nil
}

// old solution
func firstPart(fPairs []Pair, sPairs [][]int) int {
	midTotal := 0
	for _, sPair := range sPairs {
		if isValidPair(fPairs, sPair) {
			mid := sPair[len(sPair)/2]
			midTotal += mid
		}
	}
	return midTotal
}

// old solution
func isValidPair(fPairs []Pair, sPair []int) bool {
	var isValid bool
	// prepare pairs
	p := [][]int{}
	for id := range sPair {
		for j := id + 1; j < len(sPair); j++ {
			p = append(p, []int{sPair[id], sPair[j]})
		}
	}

	for _, pair := range p {
		temp := false
		for _, fPair := range fPairs {
			if fPair.First == pair[0] && fPair.Second == pair[1] {
				temp = true
				break
			}
		}
		if !temp {
			isValid = false
			break
		}
		isValid = true
	}

	return isValid
}
