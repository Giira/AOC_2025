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

func biggestRectangle(coords []Coord) ([]Rectangle, Coord) {
	var rectangles []Rectangle
	var maxX int
	var maxY int
	for i := 0; i < len(coords); i++ {
		maxX = max(maxX, coords[i].x)
		maxY = max(maxY, coords[i].y)
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
	return rectangles, Coord{x: maxX, y: maxY}
}

type strCoord struct {
	x   int
	y   int
	str string
}

func makeGrid(min Coord, max Coord) [][]strCoord {
	var out [][]strCoord
	for i := min.x; i <= max.x; i++ {
		var outInner []strCoord
		for j := min.y; j <= max.y; j++ {
			outInner = append(outInner, strCoord{x: i, y: j, str: "."})
		}
		out = append(out, outInner)
	}
	return out
}

func makeShape(grid [][]strCoord, coords []Coord) {
	coords = append(coords, coords[0])
	for i := 0; i < (len(coords) - 1); i++ {

	}
}

func day9() {
	data := import_file("day9.txt")
	coords := lines_to_coords(data)
	rectangles, maxs := biggestRectangle(coords)
	grid := makeGrid(Coord{x: 0, y: 0}, maxs)
	fmt.Printf("Day 9, Part 1: %v\n", rectangles[0].area)
	fmt.Println(grid)
}
