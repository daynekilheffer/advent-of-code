package main

import (
	"fmt"
	"os"
	"strconv"
)

type dial struct {
	position int
	tick     int
}

func (d *dial) Rotate(steps int) {
	absSteps := steps
	step := 1
	if steps < 0 {
		absSteps = -steps
		step = -1
	}
	for i := 0; i < absSteps; i++ {
		d.position = (100 + d.position + step) % 100
		if d.position == 0 {
			d.tick++
		}
	}
}

func main() {
	d := dial{position: 50}

	input := os.Args[1]
	f, err := os.OpenFile(input, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

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
		iter := steps
		if line[0] == 'L' {
			iter = -steps
		}
		d.Rotate(iter)
	}
	fmt.Printf("Final position: %d\n", d.position)
	fmt.Printf("Number of times at position 0: %d\n", d.tick)
}
