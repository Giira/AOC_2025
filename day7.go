package main

import (
	"fmt"
)

func fire_tachyon(tachyon_map map[Coord]string, data []string) (map[Coord]string, int, int) {
	counter := 0
	beam_sums := make(map[int]int, 0)
	for i := range len(data) - 1 {
		for j := range len(data[0]) {
			c := Coord{
				x: i,
				y: j,
			}
			above := Coord{
				x: i - 1,
				y: j,
			}
			if tachyon_map[c] == "S" {
				c.x += 1
				tachyon_map[c] = "|"
				beam_sums[c.y] = 1
			} else if tachyon_map[c] == "^" && tachyon_map[above] == "|" {
				beam_sums[c.y-1] += beam_sums[c.y]
				beam_sums[c.y+1] += beam_sums[c.y]
				delete(beam_sums, c.y)
				counter += 1
				c.y -= 1
				tachyon_map[c] = "|"
				c.y += 2
				tachyon_map[c] = "|"
			} else if tachyon_map[c] == "|" {
				c.x += 1
				if tachyon_map[c] != "^" {
					tachyon_map[c] = "|"
				}
			} else if tachyon_map[c] == "." && tachyon_map[above] == "|" {
				tachyon_map[c] = "|"
			}
		}
	}
	beam_sum := 0
	for _, k := range beam_sums {
		beam_sum += k
	}
	return tachyon_map, counter, beam_sum
}

func day7() {
	data := import_file("day7.txt")
	tachyon_map := split_chars_coord_map(data)
	_, count, sum := fire_tachyon(tachyon_map, data)
	fmt.Printf("Day 7, Part 1: %v\n", count)
	fmt.Printf("Day 7, Part 2: %v\n", sum)
}
