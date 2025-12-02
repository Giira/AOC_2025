package main

import (
	"bufio"
	"log"
	"os"
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
