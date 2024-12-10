package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

const day = "day-4"
const fileName = day + "-input.txt"

type Pair struct {
	a int
	b int
}

func main() {

	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	lines, err := getHLines(file)
	if err != nil {
		log.Fatal(err)
	}

	part1 := func() int {
		checkArm := func(y int, x int, dy int, dx int) bool {
			// log.Println("checkArm:", y, x, dy, dx)

			pattern := "MAS"
			for n := range 3 { // Range is 0-2
				newY := y + (dy * (n + 1))
				newX := x + (dx * (n + 1))
				// log.Println("CHECKING:", newY, newX)

				if newY < 0 || newY >= len(lines) || newX < 0 || newX >= len(lines[0]) {
					return false
				}

				// log.Println("TEST:", string(lines[newY][newX]), string(pattern[n]))
				if lines[newY][newX] != pattern[n] {
					return false
				}
			}

			// log.Println("FOUND:", y, x)
			return true
		}

		good := 0
		for y, line := range lines {
			for x, char := range line {
				// log.Println(y, x, string(char))
				if string(char) == "X" {
					for _, dy := range [3]int{-1, 0, 1} {
						for _, dx := range [3]int{-1, 0, 1} {
							if checkArm(y, x, dy, dx) {
								good += 1
							}
						}
					}
				}
			}
		}

		return good
	}

	part2 := func() int {
		checkDiags := func(y int, x int) bool {
			// log.Println("checkDiags:", y, x)
			// Make sure the middle is within 1 space of edges, since the X needs one space to each direction.
			if y < 1 || y >= (len(lines)-1) || x < 1 || x >= (len(lines[0])-1) {
				// log.Println("Out of bounds.")
				return false
			}

			uldr := string(lines[y-1][x-1]) + string(lines[y+1][x+1]) // up-left to down-right
			// log.Println("TESTING:", uldr)
			if !(uldr == "MS" || uldr == "SM") {
				// log.Println("NOPE")
				return false
			}

			urdl := string(lines[y-1][x+1]) + string(lines[y+1][x-1]) // up-right to down-left
			// log.Println("TESTING:", uldr)
			if !(urdl == "MS" || urdl == "SM") {
				// log.Println("NOPE")
				return false
			}

			return true
		}

		good := 0
		for y, line := range lines {
			for x, char := range line {
				// log.Println(y, x, string(char))
				if string(char) == "A" {
					if checkDiags(y, x) {
						good += 1
					}
				}
			}
		}

		return good
	}
	fmt.Println("[Part 1] Total = ", part1())
	fmt.Println("[Part 2] Total:", part2())
}

// getHLines returns horizontal lines array
func getHLines(file io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// getVLines returns vertical lines array from horizontal lines
func getVLines(lines []string) ([]string, error) {

	if len(lines) == 0 {
		return nil, fmt.Errorf("empty slice")
	}

	rows, cols := len(lines), len(lines[0])
	columns := make([]string, cols)

	for col := range cols {
		for row := range rows {
			columns[col] += string(lines[row][col])
		}
	}

	return columns, nil
}

func getHCount(lines []string) (int, error) {
	var count int

	re, err := regexp.Compile(`(?m)XMAS|SAMX`)
	if err != nil {
		return count, err
	}

	for _, line := range lines {
		count += len(re.FindAllString(line, -1))
	}

	return count, nil
}

// extractDiagonals extracts all diagonals: top-left to bottom-right, top-right to bottom-left,
// bottom-left to top-right, and bottom-right to top-left.
func extractDiagonals(grid []string) ([]string, []string, []string, []string) {
	if len(grid) == 0 {
		return nil, nil, nil, nil
	}
	rows, cols := len(grid), len(grid[0])
	tlBr := []string{} // Top-left to bottom-right
	trBl := []string{} // Top-right to bottom-left
	blTr := []string{} // Bottom-left to top-right
	brTl := []string{} // Bottom-right to top-left

	// Top-left to bottom-right diagonals
	for d := 0; d < rows+cols-1; d++ {
		var diag string
		for row := 0; row < rows; row++ {
			col := d - row
			if col >= 0 && col < cols {
				diag += string(grid[row][col])
			}
		}
		if len(diag) > 0 {
			tlBr = append(tlBr, diag)
		}
	}

	// Top-right to bottom-left diagonals
	for d := 0; d < rows+cols-1; d++ {
		var diag string
		for row := 0; row < rows; row++ {
			col := cols - 1 - (d - row)
			if col >= 0 && col < cols {
				diag += string(grid[row][col])
			}
		}
		if len(diag) > 0 {
			trBl = append(trBl, diag)
		}
	}

	// Bottom-left to top-right diagonals
	for d := 0; d < rows+cols-1; d++ {
		var diag string
		for row := rows - 1; row >= 0; row-- {
			col := d - (rows - 1 - row)
			if col >= 0 && col < cols {
				diag += string(grid[row][col])
			}
		}
		if len(diag) > 0 {
			blTr = append(blTr, diag)
		}
	}

	// Bottom-right to top-left diagonals
	for d := 0; d < rows+cols-1; d++ {
		var diag string
		for row := rows - 1; row >= 0; row-- {
			col := cols - 1 - (d - (rows - 1 - row))
			if col >= 0 && col < cols {
				diag += string(grid[row][col])
			}
		}
		if len(diag) > 0 {
			brTl = append(brTl, diag)
		}
	}

	return tlBr, trBl, blTr, brTl
}
