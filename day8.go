package main

import (
	"fmt"
	"home/euan466/adventofcode/2025/pkg/set"
	"slices"
	"sort"
)

const CONNECTIONS = 1000

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

	count := 0
	for _, distance := range distances {
		if count == limit-1 {
			break
		}
		var a int
		var b int
		for k, circuit := range circuits {
			if circuit.Contains(distance.a) {
				a = k
			}
			if circuit.Contains(distance.b) {
				b = k
			}
		}
		if a == b {
			count++
			continue
		}
		circuits[a] = circuits[a].Union(circuits[b])
		circuits = slices.Delete(circuits, b, b+1)
		count++
	}

	return circuits
}

func makeCircuitsForever(distances []Connection, CONNECTIONS int, length int) Connection {
	var circuits []*set.Set
	for i := range length {
		s := set.NewSet()
		s.Add(i)
		circuits = append(circuits, s)
	}
	var out Connection

	for _, distance := range distances {
		if len(circuits) == 1 {
			break
		}
		var a int
		var b int
		for k, circuit := range circuits {
			if circuit.Contains(distance.a) {
				a = k
			}
			if circuit.Contains(distance.b) {
				b = k
			}
		}
		if a == b {
			continue
		}
		circuits[a] = circuits[a].Union(circuits[b])
		circuits = slices.Delete(circuits, b, b+1)
		out = distance
	}
	return out
}

func part1_8(circuits []*set.Set) int {
	sizes := getSizes(circuits)
	sort.Slice(sizes, func(i, j int) bool { return sizes[i] > sizes[j] })
	out := 1
	for _, size := range sizes[:3] {
		out *= size
	}
	return out
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

func getSizes(sets []*set.Set) []int {
	var sizes []int
	for _, set := range sets {
		sizes = append(sizes, set.Size())
	}
	return sizes
}

func part2_8(distance Connection, coords []Coord3D) {
	x1 := coords[distance.a].x
	x2 := coords[distance.b].x
	out := x1 * x2
	fmt.Printf("Day 8, Part 2: %v\n", out)
}

func day8() {
	data := import_file("day8.txt")
	coords := lines_to_3d_coords(data)
	distances := calcDistances(coords)
	circuits := makeCircuits(distances, CONNECTIONS, len(coords))
	p1 := part1_8(circuits)
	fmt.Printf("Day 8, Part 1: %v\n", p1)
	d := makeCircuitsForever(distances, CONNECTIONS, len(coords))
	part2_8(d, coords)
}
