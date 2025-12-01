package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1() {
	f, err := os.Open("day1.txt")
	if err != nil {
		log.Fatalf("error: Failed to open text file: %v", err)
	}
	defer f.Close()

	var data []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	dial := 50
	counter := 0
	for _, instruction := range data {
		instruction = strings.TrimSpace(instruction)
		direction := instruction[0]
		distance_str := instruction[1:]
		if len(distance_str) != 2 && len(distance_str) != 1 {
			distance_str = distance_str[1:]
		}
		distance, err := strconv.Atoi(distance_str)
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
			dial += 100
		} else if dial > 99 {
			dial -= 100
		}
		if dial == 0 {
			counter += 1
		}
	}
	fmt.Printf("Day 1, Part 1: %v\n", counter)
}
