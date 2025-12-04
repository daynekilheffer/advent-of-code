package main

import (
	"fmt"
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
		parts := strings.Split(line, "-")
		left := parts[0]
		right := parts[1]
		fmt.Printf("Left: %s, Right: %s\n", left, right)

		invalidIds := checkRange(left, right)
		for _, id := range invalidIds {
			fmt.Println("\tinvalid ID:", id)
			total += id
		}

	}
	fmt.Printf("Total: %d\n", total)
}
func checkRange(startStr, endStr string) []int64 {
	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		panic(err)
	}
	var invalidIds []int64
	for i := start; i <= end; i++ {
		s := strconv.FormatInt(i, 10)
		if !isValid(s) {
			invalidIds = append(invalidIds, i)
		}
	}
	return invalidIds
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return true
	}
	mid := len(s) / 2
	left := s[:mid]
	right := s[mid:]
	// fmt.Println(" Comparing", left, "and", right, ":", left != right)
	return left != right
}
