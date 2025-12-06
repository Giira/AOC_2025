package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func split_input(data []string) ([]IDRange, []int) {
	ranges := true
	var out_ranges []IDRange
	var out_ids []int
	for _, line := range data {
		if line == "" {
			ranges = false
			continue
		}
		if ranges {
			tmp := strings.Split(line, "-")
			s, err := strconv.Atoi(tmp[0])
			e, err1 := strconv.Atoi(tmp[1])
			if err != nil || err1 != nil {
				log.Fatalf("error converting string to integer: %v, %v", err, err1)
			}
			out_ranges = append(out_ranges, IDRange{
				start: s,
				end:   e,
			})
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("error converting string to integer: %v", err)
			}
			out_ids = append(out_ids, i)
		}
	}
	return out_ranges, out_ids
}

func part1_5(ranges []IDRange, ids []int) {
	total := 0
	for _, id := range ids {
		for _, idrange := range ranges {
			if idrange.start <= id && id <= idrange.end {
				total += 1
				break
			}
		}
	}
	fmt.Printf("Day 5, Part 1: %v\n", total)
}

func part2_5(ranges []IDRange) {

	slices.SortFunc(ranges, sortByStart)
	merged_ranges := mergeRanges(ranges)
	total := 0
	for _, r := range merged_ranges {
		total += r.end - r.start + 1
	}

	fmt.Printf("Day 5, Part 2: %v\n", total)
}

func mergeRanges(ranges []IDRange) []IDRange {
	for {
		change := len(ranges)
		for i := 0; i < len(ranges)-1; i++ {
			for j := i + 1; j < len(ranges); j++ {
				if inRange(ranges[i], ranges[j].start) && inRange(ranges[i], ranges[j].end) {
					ranges = delRange(ranges, j)
				} else if inRange(ranges[i], ranges[j].start) && !inRange(ranges[i], ranges[j].end) {
					ranges[i].end = ranges[j].end
					ranges = delRange(ranges, j)
				}
			}
		}
		if len(ranges) == change {
			break
		}
	}
	return ranges
}

func sortByStart(a IDRange, b IDRange) int {
	if a.start < b.start {
		return -1
	} else if a.start > b.start {
		return 1
	} else {
		return 0
	}
}

func inRange(idrange IDRange, id int) bool {
	return id >= idrange.start && id <= idrange.end
}

func delRange(ranges []IDRange, idx int) []IDRange {
	return append(ranges[:idx], ranges[idx+1:]...)
}

func day5() {
	data := import_file("day5.txt")
	ranges, ids := split_input(data)
	part1_5(ranges, ids)
	part2_5(ranges)
}
