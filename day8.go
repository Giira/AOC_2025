package main

import (
	"fmt"
	"home/euan466/adventofcode/2025/pkg/set"
	"slices"
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
			continue
		}
		circuits[a] = circuits[a].Union(circuits[b])
		circuits = slices.Delete(circuits, b, b+1)
		count++
	}

	return circuits
}

func part1_8(circuits []*set.Set) int {
	circuits = slices.SortedFunc(circuits, sortSetsSize)
	out := 1
	for _, circuit := range circuits[:3] {
		out *= circuit.Size()
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

func sortSetsSize(a *set.Set, b *set.Set) int {
	if a.Size() < b.Size() {
		return -1
	} else if a.Size() > b.Size() {
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
	p1 := part1_8(circuits)
	fmt.Printf("Day 8, Part 1: %v", p1)
}
