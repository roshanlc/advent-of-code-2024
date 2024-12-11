package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	part1 := firstPart(fPairs, sPairs)
	fmt.Println("[Part 1] Total:", part1)

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

func firstPart(fPairs []Pair, sPairs [][]int) int {
	midTotal := 0
	for _, sPair := range sPairs {
		// prepare pairs
		p := [][]int{}
		for id := range sPair {
			for j := id + 1; j < len(sPair); j++ {
				p = append(p, []int{sPair[id], sPair[j]})
			}
		}
		// loop through pairs and check if pair exists or not
		exists := false
		for _, pair := range p {
			temp := false
			for _, fPair := range fPairs {
				if fPair.First == pair[0] && fPair.Second == pair[1] {
					temp = true
					break
				}
			}
			if !temp {
				exists = false
				break
			}
			exists = true
		}

		if exists {
			mid := sPair[len(sPair)/2]
			midTotal += mid
		}
	}
	return midTotal
}
