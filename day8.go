package main

import (
	"fmt"
	"home/euan466/adventofcode/2025/pkg/set"
	"slices"
)

const CONNECTIONS = 2

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

func makeCircuits(distances []Connection, CONNECTIONS int, length int) []*set.Set {
	var circuits []*set.Set
	limit := CONNECTIONS
	for i := range length {
		s := set.NewSet()
		s.Add(i)
		circuits = append(circuits, s)
	}

	for j := 1; j < limit; j++ {
		for k, circuit := range circuits {
			if k == j {
				continue
			} else if circuits[j].Intersection(circuit) != nil {
				circuits[j] = circuits[j].Union(circuit)
				circuits = slices.Delete(circuits, k, k+1)
			}
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
	circuits := makeCircuits(distances, CONNECTIONS, len(coords))
	for _, item := range circuits {
		fmt.Println(item.List())
	}

}
