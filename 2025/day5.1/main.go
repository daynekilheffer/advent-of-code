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
	for scanner.Scan() {
		itemRow := scanner.Text()
		item, err := strconv.ParseInt(itemRow, 10, 64)
		if err != nil {
			panic(err)
		}
		if freshInventory.IsItemFresh(item) {
			total++
		}
	}
	fmt.Printf("Total: %d\n", total)
}
