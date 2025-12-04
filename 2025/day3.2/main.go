package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := os.Args[1]
	f, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := int64(0)
	for {
		var line string
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err != nil {
			break
		}
		positions := make([]int64, 12)
		// fmt.Println("Processing line:", line)
		for idx, v := range line {
			val, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				panic(err)
			}
			startOffset := 0
			if len(line)-1-idx < len(positions) {
				startOffset = len(positions) - (len(line) - 1 - idx) - 1
			}
			// fmt.Printf(" checking value at %d: %d, start: %d\n", idx, val, startOffset)
			for p := startOffset; p < len(positions); p++ {
				// fmt.Printf("  is %d greater than %d\n", val, positions[p])
				if positions[p] < val {
					positions[p] = val
					for rp := p + 1; rp < len(positions); rp++ {
						positions[rp] = 0
					}
					break
				}
			}
			// fmt.Printf(" positions: %v\n", positions)
		}
		result := int64(0)
		for idx, val := range positions {
			result += val * int64(math.Pow10(len(positions)-1-idx))
		}
		total += result
		fmt.Println(strings.TrimSpace(line), "->", result)

	}
	fmt.Printf("Total: %d\n", total)
}
