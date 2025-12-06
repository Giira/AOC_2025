package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func multiply(numbers []int) int {
	total := 1
	for _, number := range numbers {
		total *= number
	}
	return total
}

func prepareForMaths(digits [][]string) [][]int {
	var out [][]int
	for i := 0; i < len(digits[0]); i++ {
		var part_out []int
		for j := 0; j < len(digits); j++ {
			num, err := strconv.Atoi(digits[j][i])
			if err != nil {
				log.Fatalf("error converting string to int: %v", err)
			}
			part_out = append(part_out, num)
		}
		out = append(out, part_out)
	}
	return out
}

func doMaths(digits [][]int, instructions []string) int {
	total := 0
	for i, instruction := range instructions {
		switch instruction {
		case "*":
			total += multiply(digits[i])
		case "+":
			total += sum(digits[i])
		}
	}
	return total
}

func day6() {
	data := import_file("day6.txt")
	digits := data[:len(data)-1]
	var dig_fields [][]string
	for _, line := range digits {
		dig_fields = append(dig_fields, strings.Fields(line))
	}
	prep_digits := prepareForMaths(dig_fields)
	instructions := strings.Fields(data[len(data)-1])
	total := doMaths(prep_digits, instructions)
	fmt.Printf("Day 6, Part 1: %v\n", total)
}
