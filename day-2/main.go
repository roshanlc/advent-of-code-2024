package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const day = "day-2"
const fileName = day + "-input.txt"

/*
Part 1:
Verifying whether a report is safe or not ?
  - 1. All levels in a report should be either increasing or decreasing . ie. x < y < z or x > y > z
  - 2. The difference between adjacent levels in a report should be atleast 1 or at most 3 , i.e | x - y | >= 1 && |x -y| <= 3

Part-2:
 - The reactor can tolerate a single bad value
 - That is, if a single bad value can be removed and report can be made safe then it is valid (safe) report
*/

func main() {

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reports, err := getReports(file)
	if err != nil {
		log.Fatal(err)
	}
	safeCount, unsafeCount := processReportsWithoutDampener(reports)
	fmt.Printf("[Without Dampener] : Safe count: %d, unsafe count: %d\n", safeCount, unsafeCount)

	safeCount, unsafeCount = processReportsWithDampener(reports)
	fmt.Printf("[With Dampener] : Safe count: %d, unsafe count: %d\n", safeCount, unsafeCount)
}

// getReports extracts reports from the file content
// returns a slice (report) containing slices of levels
func getReports(file io.Reader) ([][]int, error) {
	var reports [][]int
	if file == nil {
		return nil, fmt.Errorf("could not read file: nil file found")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		parts := strings.Split(line, " ")
		levels := []int{}

		for _, level := range parts {
			l, err := strconv.Atoi(level)
			if err != nil {
				return nil, fmt.Errorf("error during conversion of string to int: %w", err)
			}
			levels = append(levels, l)
		}
		reports = append(reports, levels)
	}

	return reports, nil
}

// processReportsWithoutDampener processes reports without dampener logic
func processReportsWithoutDampener(reports [][]int) (int, int) {
	var safeCount, unsafeCount int

	for _, report := range reports {
		isValid := isValidReport(report)
		if isValid {
			safeCount++
		} else {
			unsafeCount++
		}
	}

	return safeCount, unsafeCount
}

// processReportsWithDampener processes reports with the dampener logic
func processReportsWithDampener(reports [][]int) (int, int) {
	var safeCount, unsafeCount int

	var inValidReports [][]int
	for _, report := range reports {
		isValid := isValidReport(report)
		if isValid {
			safeCount++
		} else {
			unsafeCount++
			inValidReports = append(inValidReports, report)
		}
	}

	// process invalid reports to see if dampener can work with the reports
	for _, report := range inValidReports {
		// see if report becomes valid by removing current item from slice
		for idx, _ := range report {

			var leftSlice, rightSlice, newReport []int
			if idx < len(report) {
				leftSlice = report[:idx]
				rightSlice = report[idx+1:]
			}

			newReport = append(newReport, leftSlice...)
			newReport = append(newReport, rightSlice...)

			// check if the new report is valid one
			isValid := isValidReport(newReport)
			if isValid {
				safeCount++
				unsafeCount--
				break
			}
		}
	}

	return safeCount, unsafeCount
}

// isValidReport checks the validty of the report
func isValidReport(report []int) bool {
	isValid := true

	isAscending := false
	lastLevel := 0

	for id, level := range report {
		// if the level is the first entry, assign it to lastLevel and continue
		if id == 0 {
			lastLevel = level

			// check if next item exists in slice
			// for determining order of items in report
			if len(report) > 1 {
				if (report[id+1]) > level {
					isAscending = true
				}
			}

			continue
		}

		// now for other remaining items
		if isAscending {
			// first condition
			if lastLevel > level {
				isValid = false
				break
			}

			// second condition
			diff := absDiff(lastLevel, level)
			if diff < 1 || diff > 3 {
				isValid = false
				break
			}
		} else {
			if lastLevel < level {
				isValid = false
				break
			}

			diff := absDiff(lastLevel, level)
			if diff < 1 || diff > 3 {
				isValid = false
				break
			}
		}
		lastLevel = level
	}

	return isValid
}

// absDiff returns absolute difference
func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
