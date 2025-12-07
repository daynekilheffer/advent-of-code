package main

import (
	"bufio"
	"fmt"
	"os"
)

type BeamRun struct {
	Row           int
	Col           int
	TimeLineCount int
}

func (b BeamRun) Key() string {
	return fmt.Sprintf("(%d,%d)", b.Row, b.Col)
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	rowIdx := 0
	run := &BeamRun{
		TimeLineCount: 1,
	}
	splitters := make(map[string]bool)
	width := 0
	for scanner.Scan() {
		row := scanner.Bytes()
		if rowIdx == 0 {
			width = len(row)
		}
		for colIdx, char := range row {
			if char == '^' {
				splitters[fmt.Sprintf("%d,%d", rowIdx, colIdx)] = true
			}
			if rowIdx == 0 && char == 'S' {
				run.Col = colIdx
			}
		}
		rowIdx++
	}
	height := rowIdx

	referenceSplits := map[string]*BeamRun{}

	// fmt.Printf("%+v\n", run)

	queue := []*BeamRun{run}
	timelines := 1
	iterations := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		iterations++
		if iterations > 10000000000 {
			fmt.Println("exhausted iterations")
			break
		}
		// printGrid(splitters, height, width, current)
		// fmt.Println("----")

		nextRow := current.Row + 1
		if nextRow >= height {
			continue
		}
		nextPosition := fmt.Sprintf("%d,%d", nextRow, current.Col)
		if !splitters[nextPosition] {
			current.Row++
			queue = append(queue, current)
		} else {
			left := &BeamRun{
				Row: nextRow,
				Col: current.Col - 1,
			}
			right := &BeamRun{
				Row: nextRow,
				Col: current.Col + 1,
			}
			split := false
			if left.Col >= 0 && left.Col < width && left.Row < height {
				if referenceSplits[left.Key()] != nil {
					referenceSplits[left.Key()].TimeLineCount++
				} else {
					referenceSplits[left.Key()] = left
					queue = append(queue, left)
					split = true
				}
			}
			if right.Col >= 0 && right.Col < width && right.Row < height {
				if referenceSplits[right.Key()] != nil {
					referenceSplits[right.Key()].TimeLineCount++
				} else {
					referenceSplits[right.Key()] = right
					queue = append(queue, right)
					split = true
				}
			}
			if split {
				timelines += current.TimeLineCount
			}
		}

	}
	fmt.Println("total timelines", timelines)
}

func printGrid(splitters map[string]bool, height, width int, run BeamRun) {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if splitters[fmt.Sprintf("%d,%d", r, c)] {
				fmt.Print("^")
			} else if r == run.Row && c == run.Col {
				fmt.Print("x")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func printQueue(queue []BeamRun) {
	fmt.Print("Queue: ")
	for _, run := range queue {
		fmt.Printf("(%d,%d) ", run.Row, run.Col)
	}
	fmt.Println()
}
