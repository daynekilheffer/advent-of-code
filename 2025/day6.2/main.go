package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	values []string
	op     operation
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		rows = append(rows, row)
	}
	operationRow := rows[len(rows)-1]
	operationFields := regexp.MustCompile(`[*+]\s+`).FindAllString(operationRow, -1)
	fieldLengths := make([]int, len(operationFields))
	for i, field := range operationFields {
		fieldLengths[i] = len(field)
	}

	problems := make([]MathProblem, len(operationFields))
	for _, row := range rows {

		fields := extractRowToFields(row, fieldLengths)
		for _, v := range fields {
			fmt.Printf("%q ", v)
		}
		fmt.Println()

		if isOperationRow(row) {
			for idx, field := range fields {
				problem := &problems[idx]
				op := strings.TrimSpace(field)
				switch op {
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
			problem.values = append(problem.values, field)
		}
	}

	total := int64(0)
	var results []int64
	for _, problem := range problems {
		result := problem.op(convertValuesToInt64(problem.values))
		results = append(results, result)
		total += result
	}

	for _, result := range results {
		println(result)
	}
	println("Total:", total)
}

func extractRowToFields(row string, fieldLengths []int) []string {
	var fields []string
	currentIndex := 0
	for _, length := range fieldLengths {
		if currentIndex+length > len(row) {
			fields = append(fields, row[currentIndex:])
			break
		}
		field := row[currentIndex : currentIndex+length]
		fields = append(fields, field)
		currentIndex += length
	}
	return fields
}

func convertValuesToInt64(values []string) []int64 {
	var numbers []int64
	longestWord := slices.MaxFunc(values, func(a, b string) int { return len(a) - len(b) })
	longestWordLength := len(longestWord)
	fmt.Println("values", values)

	extractedValues := make([]string, 0, len(values[0]))
	for i := range longestWordLength {
		var sb strings.Builder
		for _, value := range values {
			if i < len(value) && value[i] != ' ' {
				sb.WriteByte(value[i])
			}
		}
		if sb.Len() == 0 {
			continue
		}
		extractedValues = append(extractedValues, sb.String())
	}
	fmt.Println("extractedValues", extractedValues)
	for _, value := range extractedValues {
		num, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func isOperationRow(row string) bool {
	return row[0] == '+' || row[0] == '*'
}
