package day06

import (
	"fmt"
	"strings"
)

// PlanetaryBody describes a gravity well
type PlanetaryBody struct {
	name           string
	parent         *PlanetaryBody
	orbitingBodies []*PlanetaryBody
}

//AddChild adds an orbiting planet to the parent.
func (parent *PlanetaryBody) AddChild(child *PlanetaryBody) {
	parent.orbitingBodies = append(parent.orbitingBodies, child)
}

// CalculateOrbits will give a total number of direct and indirect orbits.
func CalculateOrbits(input string) (totalOrbits int) {
	var centers []*PlanetaryBody
	allBodies := make(map[string]*PlanetaryBody)
	// allBodies["COM"] = &PlanetaryBody{name: "COM"}
	rawOrbits := strings.Split(input, "\n")
	for _, rawOrbit := range rawOrbits {
		bodies := strings.Split(rawOrbit, ")")
		// var parentBody *PlanetaryBody

		if allBodies[bodies[0]] == nil {
			parentBody := PlanetaryBody{name: bodies[0]}
			allBodies[bodies[0]] = &parentBody
			centers = append(centers, &parentBody)
		}
		parentBody := allBodies[bodies[0]]
		childBody := PlanetaryBody{name: bodies[1], parent: parentBody}
		parentBody.AddChild(&childBody)
		allBodies[bodies[1]] = &childBody
	}
	fmt.Println(len(centers))
	for _, center := range centers {
		totalOrbits += countOrbits(center, 0)
	}
	return totalOrbits
}

func countOrbits(body *PlanetaryBody, depth int) (total int) {
	total += depth
	depth ++
	for _, orbitingBody := range body.orbitingBodies {
		total += countOrbits(orbitingBody, depth)
	}
	return total
}
