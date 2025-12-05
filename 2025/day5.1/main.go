package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range struct {
	start int64
	end   int64
}

func (r Range) Contains(value int64) bool {
	return value >= r.start && value <= r.end
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
	for scanner.Scan() {
		row := scanner.Text()
		_ = row
	}
	fmt.Println("Final rows:")
	fmt.Printf("Total: %d\n", total)
}
