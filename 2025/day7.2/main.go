package main

import (
	"bufio"
	"bytes"
	"fmt"
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

	rows[0] = strings.ReplaceAll(rows[0], "S", "|")

	queue := []BeamRun{
		{Grid: rows, Row: 0, Col: strings.Index(rows[0], "|")},
	}

	timelines := 1

	maxIterations := 0
	for len(queue) > 0 {
		run := queue[0]

		if maxIterations > 5000000 {
			println("Max iterations reached, stopping to prevent infinite loop.")
			break
		}
		maxIterations++

		nextRow := run.Row + 1
		if nextRow >= len(run.Grid) {
			continue
		}
		nextChar := run.Grid[nextRow][run.Col]
		if nextChar == '.' {
			run.Grid[nextRow] = run.Grid[nextRow][:run.Col] + "|" + run.Grid[nextRow][run.Col+1:]
			run.Row = nextRow
			queue[0] = run
			continue
		}
		queue = queue[1:]
		if nextChar == '^' {
			split := false
			if run.Col > 0 {
				grid := append([]string{}, run.Grid...)
				row := bytes.Clone([]byte(grid[nextRow]))
				row[run.Col-1] = '|'
				grid[nextRow] = string(row)
				nextRun := BeamRun{Grid: grid, Row: nextRow, Col: run.Col - 1}
				queue = append(queue, nextRun)
				split = true
			}
			if run.Col < len(run.Grid[0])-1 {
				grid := append([]string{}, run.Grid...)
				row := bytes.Clone([]byte(grid[nextRow]))
				row[run.Col+1] = '|'
				grid[nextRow] = string(row)
				nextRun := BeamRun{Grid: grid, Row: nextRow, Col: run.Col + 1}
				queue = append(queue, nextRun)
				split = true
			}
			if split {
				timelines++
			}
		}
	}
	fmt.Println("total timelines", timelines)
}

func printGrid(grid []string) {
	for _, row := range grid {
		println(row)
	}
}
