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

	fmt.Println("Total difference: ", total)
	// Total difference: 1882714
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
