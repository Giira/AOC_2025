package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day1() {
	data := import_file("day1.txt")
	dial := 50
	counter := 0
	counter_2 := 0
	for _, instruction := range data {
		instruction = strings.TrimSpace(instruction)
		direction := instruction[0]
		distance_str := instruction[1:]
		add_turns_str := "0"
		if len(distance_str) != 2 && len(distance_str) != 1 {
			add_turns_str = string(distance_str[0])
			distance_str = distance_str[1:]
		}
		distance, err := strconv.Atoi(distance_str)
		if err != nil {
			log.Fatalf("error: Failed to convert distance string to integer: %v", err)
		}
		add_turns, err := strconv.Atoi(add_turns_str)
		if err != nil {
			log.Fatalf("error: Failed to convert distance string to integer: %v", err)
		}

		switch direction {
		case 'R':
			dial += distance
		case 'L':
			dial -= distance
		default:
			log.Printf("Direction not specified: %v", instruction)
		}
		if dial < 0 {
			if dial+distance != 0 {
				counter_2 += 1
			}
			dial += 100
		} else if dial > 99 {
			dial -= 100
			if dial != 0 {
				counter_2 += 1
			}
		}
		if dial == 0 && distance != 0 {
			counter += 1
			counter_2 += 1
		}
		counter_2 += add_turns
	}
	fmt.Printf("Day 1, Part 1: %v\n", counter)
	fmt.Printf("Day 1, Part 2: %v\n", counter_2)
}
