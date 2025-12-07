package main

import (
	"bufio"
	"os"
	"strings"
)

type BeamRun struct {
	Grid []string
	Row  int
	Col  int
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}

	rows[0] = strings.Replace(rows[0], "S", "|", -1)

	splitCount := 0

	for rowIdx, row := range rows {
		if rowIdx == 0 {
			continue
		}
		prevRow := rows[rowIdx-1]
		newRow := []rune(row)
		for charIdx := range len(row) {
			if newRow[charIdx] == '.' && prevRow[charIdx] == '|' {
				newRow[charIdx] = '|'
			}
			if newRow[charIdx] == '^' && prevRow[charIdx] == '|' {
				splitCount++
				if charIdx > 0 {
					newRow[charIdx-1] = '|'
				}
				if charIdx < len(row)-1 {
					newRow[charIdx+1] = '|'
				}
			}
		}
		rows[rowIdx] = string(newRow)
		printGrid(rows)
	}
	println("Total splits:", splitCount)
}

func printGrid(grid []string) {
	for _, row := range grid {
		println(row)
	}
	println()
}
