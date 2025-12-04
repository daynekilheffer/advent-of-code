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
		tens := int64(-1)
		ones := int64(-1)
		for idx, v := range line {
			val, err := strconv.ParseInt(string(v), 10, 64)
			if err != nil {
				panic(err)
			}
			if tens < val && idx < len(line)-1 {
				tens = val
				ones = -1
				continue
			}
			if ones < val {
				ones = val
			}
		}
		total += tens*10 + ones
		fmt.Println(strings.TrimSpace(line), "->", tens*10+ones)

	}
	fmt.Printf("Total: %d\n", total)
}
