package main

import (
	"fmt"
	"os"
	"strconv"
)

type dial struct {
	position int
}

func (d *dial) Rotate(steps int) {
	d.position = (100 + d.position + steps) % 100
}

func main() {
	d := dial{position: 50}

	input := os.Args[1]
	f, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	zeroCount := 0
	for {
		var line string
		_, err := fmt.Fscanf(f, "%s\n", &line)
		if err != nil {
			break
		}
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch line[0] {
		case 'L':
			d.Rotate(-steps)
		case 'R':
			d.Rotate(steps)
		}
		if d.position == 0 {
			zeroCount++
		}
	}
	fmt.Printf("Final position: %d\n", d.position)
	fmt.Printf("Number of times at position 0: %d\n", zeroCount)
}
