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
	ilist := *inv
	slices.SortStableFunc(ilist, func(a, b InventoryRange) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		}
		return 0
	})
	for i := range ilist {
		newStart := ilist[i].start
		// if any previous range contains this start, skip to the end of that range
		for j := range i {
			if ilist[j].Contains(newStart) {
				newStart = ilist[j].end + 1
			}
		}
		if newStart > ilist[i].end {
			continue
		}
		newEnd := ilist[i].end
		// if any previous range contains this end, shorten to the start of that range
		for j := range i {
			if ilist[j].Contains(newEnd) {
				newEnd = ilist[j].start - 1
			}
		}
		if newEnd < newStart {
			continue
		}
		count += (newEnd - newStart + 1)
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
