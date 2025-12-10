package main

import (
	"fmt"
)

func fire_tachyon(tachyon_map map[Coord]string, data []string) (map[Coord]string, int) {
	counter := 0
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
			} else if tachyon_map[c] == "^" && tachyon_map[above] == "|" {
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
	return tachyon_map, counter
}

func day7() {
	data := import_file("day7.txt")
	tachyon_map := split_chars_coord_map(data)
	_, count := fire_tachyon(tachyon_map, data)
	fmt.Printf("Day 7, Part 1: %v\n", count)
}
