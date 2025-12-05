package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type InventoryRange struct {
	start int64
	end   int64
}

func (r InventoryRange) Contains(value int64) bool {
	return value >= r.start && value <= r.end
}

type InventoryList []InventoryRange

func (inv *InventoryList) IsItemFresh(item int64) bool {
	if inv == nil {
		return false
	}
	return slices.ContainsFunc(*inv, func(ir InventoryRange) bool {
		return ir.Contains(item)
	})
}

func (inv *InventoryList) CountUniqueItems(item int64) int64 {
	if inv == nil {
		return 0
	}
	count := int64(0)
	seenIds := make(map[int64]struct{})
	for _, invRange := range *inv {
		for i := invRange.start; i <= invRange.end; i++ {
			if _, ok := seenIds[i]; !ok {
				seenIds[i] = struct{}{}
				count++
			}
		}
	}
	return count
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
	freshInventory := InventoryList{}
	for scanner.Scan() {
		inventoryRow := scanner.Text()
		if inventoryRow == "" {
			break
		}

		inventoryParts := strings.SplitN(inventoryRow, "-", 2)
		start, err := strconv.ParseInt(inventoryParts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		end, err := strconv.ParseInt(inventoryParts[1], 10, 64)
		if err != nil {
			panic(err)
		}
		freshInventory = append(freshInventory, InventoryRange{
			start: start,
			end:   end,
		})
	}
	total = freshInventory.CountUniqueItems(0)
	fmt.Printf("Total: %d\n", total)
}
