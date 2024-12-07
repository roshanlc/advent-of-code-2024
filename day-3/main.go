package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const day = "day-3"
const fileName = day + "-input.txt"

type Pair struct {
	a int
	b int
}

func main() {

	// open file
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	pairs, err := extractMatchingExpr(&fileContents)

	if err != nil {
		log.Fatal(err)
	}

	totalPairMultSum := addAllPairs(pairs)
	fmt.Println("[Part 1]Total :", totalPairMultSum)

	// part two
	pairs, err = extractWithConsideration(&fileContents)
	if err != nil {
		log.Fatal(err)
	}
	totalPairMultSum = addAllPairs(pairs)
	fmt.Println("[Part 2]Total :", totalPairMultSum)
}

func extractMatchingExpr(contents *[]byte) ([]Pair, error) {
	re, err := regexp.Compile(`(?m)(\B|\B)?(mul\(\d{1,3},\d{1,3}\))(\B|\B)?`)

	if err != nil {
		return nil, err
	}

	reNum, err := regexp.Compile(`(?m)\d+`)

	if err != nil {
		return nil, err
	}

	var pairs []Pair
	for _, match := range re.FindAllString(string(*contents), -1) {
		var a, b int
		nums := reNum.FindAllString(match, -1)
		if len(nums) == 0 || len(nums) > 2 {
			continue
		}

		a, _ = strconv.Atoi(nums[0])
		b, _ = strconv.Atoi(nums[1])

		p := Pair{
			a: a,
			b: b,
		}

		pairs = append(pairs, p)
	}
	return pairs, nil
}

func addAllPairs(pairs []Pair) int {
	var total int

	for _, p := range pairs {
		mult := p.a * p.b
		total += mult
	}
	return total
}

func extractWithConsideration(contents *[]byte) ([]Pair, error) {

	strData := string(*contents)
	re, err := regexp.Compile(`(?m)((?:don't|do|mul)\(\d*,?\d*\))`)
	if err != nil {
		log.Fatal(err)
	}
	reNum, err := regexp.Compile(`(?m)\d+`)

	if err != nil {
		return nil, err
	}

	var pairs []Pair
	enabled := true
	for _, match := range re.FindAllString(strData, -1) {

		if match == "don't()" {
			enabled = false
		} else if match == "do()" {
			enabled = true
		} else {
			if enabled {
				var a, b int
				nums := reNum.FindAllString(match, -1)
				if len(nums) != 2 {
					continue
				}
				a, _ = strconv.Atoi(nums[0])
				b, _ = strconv.Atoi(nums[1])
				p := Pair{
					a: a,
					b: b,
				}
				pairs = append(pairs, p)
			}
		}
	}
	return pairs, nil
}
