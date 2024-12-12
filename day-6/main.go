package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const day = "day-6"
const fileName = day + "-input.txt"

type Position struct {
	X int
	Y int
}

func main() {
	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := getLines(file)
	if err != nil {
		log.Fatal(err)
	}

	part1 := part1(lines)

	fmt.Println("[Part 1] Total:", part1)
	// fmt.Println("[Part 2] Total:", part2)
}

func part1(lines []string) int {
	var guardChar = "^" // starting sign , other signs: [^, v, >, <]

	var currentPos []int
	// find guard position
	for idx, line := range lines {
		if strings.Contains(line, guardChar) {
			currentPos = []int{idx, strings.Index(line, guardChar)}
			break
		}
	}

	if currentPos == nil {
		return -1
	}

	fmt.Println("Starting positon:", currentPos)
	foundExit := false
	// until the end is not found
	var visited = make(map[Position]bool)
	visited[Position{
		X: currentPos[0],
		Y: currentPos[1],
	}] = true
	// modify the starting positon line
	l := []rune(lines[currentPos[0]])
	l[currentPos[1]] = '.'
	lines[currentPos[0]] = string(l)

	for id, l := range lines {
		fmt.Println(id+1, l)
	}
	for !foundExit {
		if (currentPos[0] == 0 ||
			currentPos[0] == (len(lines)-1) ||
			currentPos[1] == 0 ||
			currentPos[1] == (len(lines[0])-1)) &&
			string(lines[currentPos[0]][currentPos[1]]) != "#" {
			foundExit = true
			break
		}
		// current position should always be a valid position
		var temp []int
		copy(temp, currentPos) // copy the current poisition
		if guardChar == "^" {
			x := currentPos[0] - 1
			y := currentPos[1]
			if string(lines[x][y]) != "#" {
				currentPos[0] -= 1

			} else {
				guardChar = ">"
				currentPos[1] += 1
			}
			visited[Position{
				X: currentPos[0],
				Y: currentPos[1],
			}] = true
			continue
		} else if guardChar == "v" {
			if string(lines[currentPos[0]+1][currentPos[1]]) != "#" {
				currentPos[0] += 1

			} else {
				guardChar = "<"
				currentPos[1] -= 1
			}
			visited[Position{
				X: currentPos[0],
				Y: currentPos[1],
			}] = true
			continue
		} else if guardChar == ">" {
			if string(lines[currentPos[0]][currentPos[1]+1]) != "#" {
				currentPos[1] += 1

			} else {
				guardChar = "v"
				currentPos[0] += 1
			}
			visited[Position{
				X: currentPos[0],
				Y: currentPos[1],
			}] = true
			continue

		} else if guardChar == "<" {
			if string(lines[currentPos[0]][currentPos[1]-1]) != "#" {
				currentPos[1] -= 1

			} else {
				guardChar = "^"
				currentPos[0] -= 1
			}
			visited[Position{
				X: currentPos[0],
				Y: currentPos[1],
			}] = true
			continue
		}
	}

	return len(visited)
}

func getLines(reader io.Reader) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, text)
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}
