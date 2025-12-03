package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

type Joltage struct {
	first  int
	second int
}

func make_bank(data_line string) []int {
	line := strings.Split(data_line, "")
	var bank []int
	for _, char := range line {
		i, err := strconv.Atoi(char)
		if err != nil {
			log.Fatalf("error converting char to int: %v", err)
		}
		bank = append(bank, i)
	}
	return bank
}

func part1_3(line string, result *uint64, wg *sync.WaitGroup, mu *sync.Mutex) {
	bank := make_bank(line)
	j := Joltage{
		first:  0,
		second: 0,
	}
	for i, char := range bank {
		if char > j.first && (i+1) < len(bank) {
			j.first = char
			j.second = 0
		} else if char > j.second {
			j.second = char
		}
	}
	out := fmt.Sprintf("%v%v", j.first, j.second)
	out_int, err := strconv.Atoi(out)
	if err != nil {
		log.Fatalf("error converting sprint to int")
	}
	mu.Lock()
	*result += uint64(out_int)
	mu.Unlock()

	defer wg.Done()
}

type BigJoltage struct {
	first    int
	second   int
	third    int
	fourth   int
	fifth    int
	sixth    int
	seventh  int
	eighth   int
	ninth    int
	tenth    int
	eleventh int
	twelfth  int
}

func part2_3(line string, result *uint64, wg *sync.WaitGroup, mu *sync.Mutex) {
	bank := make_bank(line)
	length := len(bank)

}

func part1_parent_3(data []string) {
	var result uint64 = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, line := range data {
		wg.Add(1)
		go part1_3(line, &result, &wg, &mu)
	}
	wg.Wait()
	fmt.Printf("Day 3, Part 1: %v\n", result)
}

func part2_parent_3(data []string) {
	var result uint64 = 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, line := range data {
		wg.Add(1)
		go part2_3(line, &result, &wg, &mu)
	}
	wg.Wait()
	fmt.Printf("Day 3, Part 2: %v\n", result)
}

func day3() {
	data := import_file("day3.txt")

	part1_parent_3(data)
}
