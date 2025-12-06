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

type IDRange struct {
	start int
	end   int
}
