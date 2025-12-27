package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

type Coord3D struct {
	x int
	y int
	z int
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

func lines_to_3d_coords(data []string) []Coord3D {
	var out []Coord3D
	for _, line := range data {
		tmp := strings.Split(line, ",")
		x, err0 := strconv.Atoi(tmp[0])
		y, err1 := strconv.Atoi(tmp[1])
		z, err2 := strconv.Atoi(tmp[2])
		if err0 != nil || err1 != nil || err2 != nil {
			log.Fatalf("error converting coordinate to int: x:%v y:%v z:%v", err0, err1, err2)
		}
		coord := Coord3D{
			x: x,
			y: y,
			z: z,
		}
		out = append(out, coord)
	}
	return out
}

func lines_to_coords(data []string) []Coord {
	var out []Coord
	for _, line := range data {
		tmp := strings.Split(line, ",")
		x, err0 := strconv.Atoi(tmp[0])
		y, err1 := strconv.Atoi(tmp[1])
		if err0 != nil || err1 != nil {
			log.Fatalf("error converting coordinate to int: x:%v, y:%v", err0, err1)
		}
		coord := Coord{
			x: x,
			y: y,
		}
		out = append(out, coord)
	}
	return out
}

type IDRange struct {
	start int
	end   int
}

type Connection struct {
	a    int
	b    int
	dist int
}
