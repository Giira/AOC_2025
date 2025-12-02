package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

type IDRange struct {
	start int
	end   int
}

func make_ranges(instructions []string) []IDRange {
	ranges := []IDRange{}
	for _, instruction := range instructions {
		pair := strings.Split(instruction, "-")
		startID, err := strconv.Atoi(pair[0])
		if err != nil {
			log.Fatalf("Error converting start string to int: %v", err)
		}
		finishID, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Fatalf("Error converting end string to int: %v", err)
		}
		ranges = append(ranges, IDRange{
			start: startID,
			end:   finishID,
		})
	}
	return ranges
}

func part1(IDrange IDRange, result *uint64, wg *sync.WaitGroup, mu *sync.Mutex) {
	for current := (IDrange.start); current <= IDrange.end; current++ {
		current_str := strconv.Itoa(current)
		if len(current_str)%2 == 0 {
			half1 := current_str[0:(len(current_str) / 2)]
			half2 := current_str[(len(current_str) / 2):]
			if half1 == half2 {
				mu.Lock()
				*result += uint64(current)
				mu.Unlock()
			}
		}
	}
	defer wg.Done()
}

func day2() {
	data := import_file("day2.txt")
	instructions := strings.Split(data[0], ",")
	ranges := make_ranges(instructions)

	var result1 uint64 = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, IDRange := range ranges {
		wg.Add(1)
		go part1(IDRange, &result1, &wg, &mu)
	}
	wg.Wait()
	fmt.Printf("Day 2, Part 1: %v\n", result1)
}
