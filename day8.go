package main

import (
	"fmt"
	"home/euan466/adventofcode/2025/pkg/set"
	"slices"
	"strconv"
)

const CONNECTIONS = 100

func square(a int) int {
	return a * a
}

func distance(a Coord3D, b Coord3D) int {
	return square(a.x-b.x) + square(a.y-b.y) + square(a.z-b.z)
}

func calcDistances(coords []Coord3D) []Connection {
	connections := []Connection{}
	nCoords := len(coords)
	for i := range coords {
		for j := i + 1; j < nCoords; j++ {
			con := Connection{
				a:    i,
				b:    j,
				dist: distance(coords[i], coords[j]),
			}
			connections = append(connections, con)
		}
	}
	slices.SortFunc(connections, sortDistances)
	return connections
}

func makeCircuits(distances []Connection, CONNECTIONS int) []*set.Set {
	var circuits []*set.Set
	limit := CONNECTIONS
	if limit > len(distances) {
		limit = len(distances)
	}
	for i := 0; i < limit; i++ {
		a := strconv.Itoa(distances[i].a)
		b := strconv.Itoa(distances[i].b)
		placed := false
		for j := range circuits {
			if circuits[j].Contains(a) || circuits[j].Contains(b) {
				circuits[j].Add(a)
				circuits[j].Add(b)
				placed = true
				break
			}
		}
		if !placed {
			s := set.NewSet()
			s.Add(a)
			s.Add(b)
			circuits = append(circuits, s)
		}
	}
	return circuits
}

func sortDistances(a Connection, b Connection) int {
	if a.dist < b.dist {
		return -1
	} else if a.dist > b.dist {
		return 1
	} else {
		return 0
	}
}

func day8() {
	data := import_file("day8.txt")
	coords := lines_to_3d_coords(data)
	distances := calcDistances(coords)
	circuits := makeCircuits(distances, CONNECTIONS)
	for _, item := range circuits {
		fmt.Println(item.List())
	}

}
