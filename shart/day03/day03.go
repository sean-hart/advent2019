package main

import (
	// "fmt"
	"strconv"
	"strings"
)

// Coord is a combination of two integers that represent a coordiate.
type Coord struct {
	x, y int
}

// GetManhattanDistance will get the Manhattan Distance for two points
func GetManhattanDistance(inputString string) (distance int) {

	return 1
}

// FindIntersections finds all intersections for two sets of directions
func FindIntersections(inputString string) (coords []Coord) {

	return []Coord{Coord{1, 2}}
}

//ExpandRoute will give all of the points on the grid.
func ExpandRoute(inputString string) []Coord {
	instructions := strings.Split(inputString, ",")
	var origin = Coord{0, 0}
	var route []Coord
	route = append(route, origin)
	current := Coord{0, 0}
	for _, inst := range instructions {
		direction := inst[:1]
		distance, _ := strconv.Atoi(inst[1:])
		// fmt.Printf("Dir: %s, Dist: %d\n", direction, distance)
		// fmt.Printf("Currnt Coord %v", current)
		switch direction {
		case "R":
			for i := current.x + 1; i <= current.x+distance; i++ {
				route = append(route, Coord{i, current.y})
			}
		case "L":
			for i := current.x - 1; i >= current.x-distance; i-- {
				route = append(route, Coord{i, current.y})
			}
		case "U":
			for i := current.y + 1; i <= current.y+distance; i++ {
				route = append(route, Coord{current.x, i})
			}
		case "D":
			for i := current.y - 1; i >= current.y-distance; i-- {
				route = append(route, Coord{current.x, i})
			}

		}
		current = route[len(route)-1]
	}
	// fmt.Println(route)
	return route
}
