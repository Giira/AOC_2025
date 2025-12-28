package main

import (
	"fmt"
	"slices"
)

type Rectangle struct {
	corner1 Coord
	corner2 Coord
	area    int
}

func makeRectangle(coord1 Coord, coord2 Coord) int {
	x := coord1.x - coord2.x
	x = max(x, -x)
	x++
	y := coord1.y - coord2.y
	y = max(y, -y)
	y++
	return x * y
}

func sortRectangles(rec1 Rectangle, rec2 Rectangle) int {
	if rec1.area < rec2.area {
		return 1
	} else if rec1.area > rec2.area {
		return -1
	} else {
		return 0
	}
}

func biggestRectangle(coords []Coord) []Rectangle {
	var rectangles []Rectangle
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			a := makeRectangle(coords[i], coords[j])
			rec := Rectangle{
				corner1: coords[i],
				corner2: coords[j],
				area:    a,
			}
			rectangles = append(rectangles, rec)
		}
	}
	slices.SortFunc(rectangles, sortRectangles)
	return rectangles
}

func day9() {
	data := import_file("day9.txt")
	coords := lines_to_coords(data)
	rectangles := biggestRectangle(coords)
	fmt.Printf("Day 9, Part 1: %v\n", rectangles[0].area)
}
