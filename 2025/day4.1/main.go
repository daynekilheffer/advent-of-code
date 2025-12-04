package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var paper byte = '@'

const distance = 1
const activeRowCount = distance*2 + 1
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := int64(0)
	activeRows := make([]GridRow, activeRowCount)
	var result []GridRow
	for scanner.Scan() {
		row := scanner.Bytes()
		activeRows = append(activeRows, bytes.Clone(row))
		if len(activeRows) > activeRowCount {
			activeRows = activeRows[1:]
		}
		// fmt.Println("active rows")
		// for _, v := range activeRows {
		// 	fmt.Println(" ", v.String())
		// }

		count, resultRow := countRollsInRow(activeRows, distance)
		result = append(result, resultRow)
		fmt.Println("row", activeRows[1].String(), "->", count)
		total += count
	}
	activeRows = append(activeRows[1:], nil)
	count, resultRow := countRollsInRow(activeRows, distance)
	result = append(result, resultRow)
	fmt.Println("row", activeRows[1].String(), "->", count)
	total += count
	fmt.Println("Final rows:")
	for _, v := range result {
		fmt.Println(" ", v.String())
	}
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
		// fmt.Println("checking", positionValue, "at", middleIterIdx)
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
				// fmt.Println("  checking", colIdx, "in", rowIdx)
				if row.isPaper(colIdx) {
					// fmt.Println("   found")
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
