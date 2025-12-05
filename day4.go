package main

import (
	"fmt"
	"strings"
)

func get_coords(coord Coord, height int, width int) []Coord {
	var out []Coord
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := coord.x + dx
			ny := coord.y + dy
			if nx >= 0 && nx < height && ny >= 0 && ny < width {
				out = append(out, Coord{x: nx, y: ny})
			}
		}
	}
	return out
}

func count_paper(data [][]string, pmap map[Coord]int) {
	height := len(data)
	if height == 0 {
		return
	}
	width := len(data[0])

	for coord := range pmap {
		if data[coord.x][coord.y] == "@" {
			total := 0
			coords := get_coords(coord, height, width)
			for _, point := range coords {
				if data[point.x][point.y] == "@" {
					total++
				}
			}
			pmap[coord] = total
		}
	}
}

func part_1_4(data [][]string, pmap map[Coord]int) {
	total := 0
	for coord := range pmap {
		if pmap[coord] < 4 && data[coord.x][coord.y] == "@" {
			total++
		}
	}
	fmt.Printf("Day 4, Part 1: %v\n", total)
}

func day4() {
	in := import_file("day4.txt")
	var data [][]string
	for _, line := range in {
		data = append(data, strings.Split(line, ""))
	}
	p_map := make_coord_map(data)
	count_paper(data, p_map)
	part_1_4(data, p_map)
}
