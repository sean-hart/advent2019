package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wires [][]string
	for scanner.Scan() {
		wire := strings.Split(scanner.Text(), ",")
		wires = append(wires, wire)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wires
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func trackWire(wire []string) [][]int {

	var track [][]int
	var position []int
	position = append(position, 0, 0)
	newPosition := append(position[:0:0], position...)
	track = append(track, newPosition)
	
	for _, val := range wire {

		direction := val[0]
		magnitude, err := strconv.Atoi(stripchars(val, "RULD"))

		if err != nil {
			fmt.Println("Error while parsing ", err)
		}

		switch direction {
		case 'R':
			position[0] += magnitude
		case 'L':
			position[0] -= magnitude
		case 'U':
			position[1] += magnitude
		case 'D':
			position[1] -= magnitude
		}

		newPosition := append(position[:0:0], position...)
		track = append(track, newPosition)
	}

	return track
}

// func findNearestIntersection(track1 [][]int, track2 [][]int) int {

// 	for
// }

func main() {
	wires := parseInput("crossedWires_input.txt")
	track1 := trackWire(wires[0])
	track2 := trackWire(wires[1])
	fmt.Println(track1)
	fmt.Println(track2)
	// fmt.Println(findNearestIntersection(track1, track2))

}
