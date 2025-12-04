package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var paper byte = '@'

const maxAdjacentPaper = 4

type GridRow []byte

func (r *GridRow) isPaper(i int) bool {
	if r == nil {
		return false
	}
	if i < 0 || i >= len(*r) {
		return false
	}
	return (*r)[i] == paper
}
func (r *GridRow) String() string {
	if r == nil {
		return "<nil>"
	}
	if len(*r) == 0 {
		return "<empty>"
	}
	return string(*r)
}

type Grid []GridRow

func (r Grid) String() string {
	var buf bytes.Buffer
	for _, row := range r {
		buf.WriteString(row.String())
		buf.WriteByte('\n')
	}
	return buf.String()
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := int64(0)
	var grid Grid
	for scanner.Scan() {
		row := scanner.Bytes()
		grid = append(grid, bytes.Clone(row))
	}
	attempts := 0
	for {
		attempts++
		if attempts > 10000 {
			panic("too many attempts")
		}
		lastTry := total
		for rowIdx, row := range grid {
			var prevRow GridRow
			if rowIdx > 0 {
				prevRow = grid[rowIdx-1]
			}
			var nextRow GridRow
			if rowIdx < len(grid)-1 {
				nextRow = grid[rowIdx+1]
			}
			activeRows := []GridRow{prevRow, row, nextRow}
			count, resultRow := countRollsInRow(activeRows, 1)
			grid[rowIdx] = resultRow
			// fmt.Println("row", row.String(), "->", count)
			total += count
		}
		if lastTry == total {
			break
		}
	}
	fmt.Println("Final rows:")
	fmt.Println(grid.String())
	fmt.Printf("Total: %d\n", total)
}

func countRollsInRow(data []GridRow, distance int) (int64, GridRow) {
	var count int64 = 0
	middleIndex := (distance - 1) + len(data) - (distance * 2)
	middle := data[middleIndex]
	matched := bytes.Clone(data[middleIndex])
	for middleIterIdx, positionValue := range middle {
		if positionValue != paper {
			continue
		}
		positionalCount := int64(0)
		for rowOffset := -distance; rowOffset <= distance; rowOffset++ {
			if positionalCount >= maxAdjacentPaper {
				break
			}
			rowIdx := middleIndex + rowOffset
			row := data[rowIdx]
			for colIdx := middleIterIdx - distance; colIdx <= (middleIterIdx + distance); colIdx++ {
				if positionalCount >= maxAdjacentPaper {
					break
				}
				// skip yourself
				if rowOffset == 0 && colIdx == middleIterIdx {
					continue
				}
				if row.isPaper(colIdx) {
					positionalCount++
				}
			}
		}
		if positionalCount < maxAdjacentPaper {
			count++
			matched[middleIterIdx] = 'x'
		}
	}
	return count, matched
}
