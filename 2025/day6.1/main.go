package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type operation func([]int64) int64

func add(values []int64) int64 {
	total := int64(0)
	for _, v := range values {
		total += v
	}
	return total
}
func multiply(values []int64) int64 {
	total := int64(1)
	for _, v := range values {
		total *= v
	}
	return total
}

type MathProblem struct {
	numbers []int64
	op      operation
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var problems []MathProblem
	for scanner.Scan() {
		row := strings.TrimSpace(scanner.Text())

		// split by any series of spaces
		fields := strings.Fields(row)
		if problems == nil {
			problems = make([]MathProblem, len(fields))
		}

		if isOperationRow(row) {
			for idx, field := range fields {
				problem := &problems[idx]
				switch field {
				case "+":
					problem.op = add
				case "*":
					problem.op = multiply
				default:
					panic("unknown operation: " + field)
				}
			}
			continue
		}

		for idx, field := range fields {
			problem := &problems[idx]
			num, err := strconv.ParseInt(field, 10, 64)
			if err != nil {
				panic(err)
			}
			problem.numbers = append(problem.numbers, num)
		}
	}

	total := int64(0)
	var results []int64
	for _, problem := range problems {
		result := problem.op(problem.numbers)
		results = append(results, result)
		total += result
	}

	for _, result := range results {
		println(result)
	}
	println("Total:", total)
}

func isOperationRow(row string) bool {
	return row[0] == '+' || row[0] == '*'
}
