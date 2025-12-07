package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	width := 0
	startCol := 0
	grid := [][]rune{}
	for scanner.Scan() {
		row := scanner.Text()
		grid = append(grid, []rune(row))
	}
	startCol = strings.Index(string(grid[0]), "S")
	width = len(grid[0])
	lookup := map[string]int{}

	for i := len(grid) - 1; i >= 0; i-- {
		for j := range width {
			val := grid[i][j]
			if val == '.' || val == 'S' {

				if i == len(grid)-1 {
					lookup[key(i, j)] = 1
				} else {
					lookup[key(i, j)] = lookup[key(i+1, j)]
				}
			}
		}
		for j := range width {
			val := grid[i][j]
			if val == '^' {

				if j > 0 {
					lookup[key(i, j)] += lookup[key(i, j-1)]
				}
				if j < width-1 {
					lookup[key(i, j)] += lookup[key(i, j+1)]
				}
			}
		}
	}

	fmt.Println(lookup[key(0, startCol)])
}

func key(i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}
