package main

import (
	// "fmt"
	"math"
	"strconv"
	"strings"
)

// Coord is a combination of two integers that represent a coordiate.
type Coord struct {
	x, y int
}

// GetManhattanDistance will get the Manhattan Distance for two points
func GetManhattanDistance(inputString string) (lowestDistance int) {
	intersections := FindIntersections(inputString)
	for _, intersection := range intersections {
		// fmt.Printf("Intersection: %v", intersection)
		tempDistance := int(math.Abs(float64(intersection.x)) + math.Abs(float64(intersection.y)))
		if lowestDistance == 0 {
			lowestDistance = tempDistance
		} else {
			if tempDistance < lowestDistance {
				lowestDistance = tempDistance
			}
		}
	}
	return int(lowestDistance)
}

// ShortestWireSum will calculate the earliest crossing of two wire paths
func ShortestWireSum(inputString string) (lowestDistance int) {
	wires := strings.Split(inputString, "\n")
	wire1coords := ExpandRoute(wires[0])
	wire2coords := ExpandRoute(wires[1])
	lowestDistance = len(wire1coords) + len(wire2coords)
	// intersections := FindIntersections(inputString)

	// use instructionset to figure out how you got to the intersection
	// maybe try closest wire1 intersection first?
	// maybe try the shortest manhattan first? -- already faster than manhattan so eff that
	// maybe try the first intersection?
	// expand route to intersections?
	// if coord1 in intersections: slice wires to coords?

	for i1, coord1 := range wire1coords {
		if i1+1 < lowestDistance {
			for i2, coord2 := range wire2coords {
				if i1+i2+2 < lowestDistance {
					if coord1 == coord2 {
						if i2 < lowestDistance {
							if i1+i2+2 < lowestDistance {
								lowestDistance = i1 + i2 + 2
							}
						}
					}
				}
			}
		}
	}
	return lowestDistance
}

// FindIntersections finds all intersections for two sets of directions
func FindIntersections(inputString string) (sharedCoords []Coord) {
	wires := strings.Split(inputString, "\n")
	wire1coords := ExpandRoute(wires[0])
	wire2coords := ExpandRoute(wires[1])

	for _, coord1 := range wire1coords {
		for _, coord2 := range wire2coords {
			if coord1 == coord2 {
				sharedCoords = append(sharedCoords, coord1)
			}
		}
	}

	return sharedCoords
}

//ExpandRoute will give all of the points on the grid.
func ExpandRoute(inputString string) (route []Coord) {
	instructions := strings.Split(inputString, ",")
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
