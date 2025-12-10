package main

import (
	"bufio"
	"log"
	"os"
)

func import_file(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("error: Failed to open text file: %v", err)
	}
	defer f.Close()

	var data []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

type Coord struct {
	x int
	y int
}

func make_coord_map(data [][]string) map[Coord]int {
	e_map := make(map[Coord]int)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			c := Coord{
				x: i,
				y: j,
			}
			e_map[c] = 0
		}
	}
	return e_map
}

func split_chars_coord_map(data []string) map[Coord]string {
	c_map := make(map[Coord]string)
	for i := range len(data) {
		for j := range len(data[i]) {
			c := Coord{
				x: i,
				y: j,
			}
			c_map[c] = string(data[i][j])
		}
	}
	return c_map
}

type IDRange struct {
	start int
	end   int
}
