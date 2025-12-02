package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type IDRange struct {
	start uint64
	end   uint64
}

func make_ranges(instructions []string) []IDRange {
	ranges := []IDRange{}
	for _, instruction := range instructions {
		pair := strings.Split(instruction, "-")
		startID, err := strconv.ParseUint(pair[0], 10, 64)
		if err != nil {
			log.Fatalf("Error converting string to uint64: %v", err)
		}
		finishID, err := strconv.ParseUint(pair[1], 10, 64)
		if err != nil {
			log.Fatalf("Error converting string to uint64: %v", err)
		}
		ranges = append(ranges, IDRange{
			start: startID,
			end:   finishID,
		})
	}
	return ranges
}

func digit_tally(digit uint64) uint {
	if digit == 0 {
		return 1
	} else {
		float_digs := math.Floor(math.Log10(float64(digit))) + 1
		return uint(float_digs)
	}

}

func day2() {
	data := import_file("day2.txt")
	instructions := strings.Split(data[0], ",")
	ranges := make_ranges(instructions)

}
