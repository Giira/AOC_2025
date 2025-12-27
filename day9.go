package main

func makeRectangle(coord1 Coord, coord2 Coord) int {
	x := coord1.x - coord2.x
	x = max(x, -x)
	y := coord1.y - coord2.y
	y = max(y, -y)
	return x * y
}

func biggestRectangle()

func day9() {
	data := import_file("day9.txt")
	coords := lines_to_coords(data)
}
