package main

import (
	"fmt"
	"home/euan466/adventofcode/2025/pkg/set"
	"slices"
	"strconv"
)

const CONNECTIONS = 10

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

func makeCircuits(distances []Connection, CONNECTIONS int) []set.Set {
	var circuits []set.Set
	for i := range CONNECTIONS {
		a := strconv.Itoa(distances[i].a)
		b := strconv.Itoa(distances[i].b)
		if len(circuits) == 0 {
			s := set.NewSet()
			s.Add(a)
			s.Add(b)
			circuits = append(circuits, *s)
		} else {
			for _, circuit := range circuits {
				if circuit.Contains(a) || circuit.Contains(b) {
					circuit.Add(a)
					circuit.Add(b)
					break
				} else {
					s := set.NewSet()
					s.Add(a)
					s.Add(b)
					circuits = append(circuits, *s)
					break
				}
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
	circuits := makeCircuits(distances, CONNECTIONS)
	fmt.Println(circuits)

}
