package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

const fileName = "day-1-input.txt"

func main() {
	// read file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	firstIDs, secondIDs, err := getLocationIDs(file)

	if err != nil {
		log.Fatal(err)
	}

	// sort the slices
	sort.Ints(firstIDs)
	sort.Ints(secondIDs)

	// run the calculation
	total := 0
	for i := 0; i < len(firstIDs); i++ {
		diff := firstIDs[i] - secondIDs[i]
		if diff < 0 {
			diff *= -1
		}
		total += diff
	}

	fmt.Println("[Part 1]Total difference: ", total)
	// Total difference: 1882714

	// part two as follows

	// sort the first location ids
	visited := map[int]bool{}

	score := 0

	for _, first := range firstIDs {

		if _, exists := visited[first]; exists {
			continue
		}

		count := 0
		for _, second := range secondIDs {
			if first == second {
				count++
			}
		}

		// add up the score
		score += (first * count)
		// set visited to true for the current item
		visited[first] = true
	}

	fmt.Println("[Part two]Score :", score)
}

// getLocationIDs scans the content of file
// for location ids and groups them into
// two sets of location IDs
func getLocationIDs(file io.Reader) ([]int, []int, error) {
	var firstIDS []int
	var secondIDS []int

	// scan tokens
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var fillFirstIDS bool = true
	for scanner.Scan() {
		item, _ := strconv.Atoi(scanner.Text())
		if fillFirstIDS {
			firstIDS = append(firstIDS, item)
		} else {
			secondIDS = append(secondIDS, item)
		}
		fillFirstIDS = !fillFirstIDS
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return firstIDS, secondIDS, nil
}
