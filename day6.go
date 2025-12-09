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

func split_numbers(in []string) [][]int {
	var out [][]int
	var out_part []int
	for _, item := range in {
		var empty []int
		if strings.Trim(item, " ") == "" {
			out = append(out, out_part)
			out_part = empty
		} else {
			o, err := strconv.Atoi(strings.Trim(item, " "))
			if err != nil {
				log.Fatalf("error converting to int: %v", err)
			}
			out_part = append(out_part, o)
		}
	}
	out = append(out, out_part)
	return out
}

func part1_6() {
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

func part2_6() {
	data := import_file("day6.txt")
	digits_in := data[:len(data)-1]
	instructions := strings.Fields(data[len(data)-1])

	var str_digits []string
	for i := range len(digits_in[0]) {
		var str_slice []string
		for j := range len(digits_in) {
			str_slice = append(str_slice, string(digits_in[j][i]))
		}
		str := strings.Join(str_slice, "")
		str_digits = append(str_digits, str)
	}
	digits_out := split_numbers(str_digits)

	total := doMaths(digits_out, instructions)
	fmt.Printf("Day 6, Part 2: %v\n", total)
}

func day6() {
	part1_6()
	part2_6()
}
