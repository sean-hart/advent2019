package day06

import (
	// "fmt"
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

func (child *PlanetaryBody) AddParent(parent *PlanetaryBody) {
	child.parent = parent
}

// CalculateOrbits will give a total number of direct and indirect orbits.
func CalculateOrbits(input string) (totalOrbits int) {
	allBodies := buildMap(input)
	totalOrbits = countOrbits(allBodies["COM"], 0)
	return totalOrbits
}

func countOrbits(body *PlanetaryBody, depth int) (total int) {
	total += depth
	depth++
	for _, orbitingBody := range body.orbitingBodies {
		total += countOrbits(orbitingBody, depth)
	}
	return total
}

func buildMap(givenMap string) (allBodies map[string]*PlanetaryBody) {
	allBodies = make(map[string]*PlanetaryBody)
	rawOrbits := strings.Split(givenMap, "\n")
	for _, rawOrbit := range rawOrbits {
		bodies := strings.Split(rawOrbit, ")")
		if allBodies[bodies[0]] == nil {
			body := PlanetaryBody{name: bodies[0]}
			allBodies[bodies[0]] = &body
		}
		if allBodies[bodies[1]] == nil {
			body := PlanetaryBody{name: bodies[1]}
			allBodies[bodies[1]] = &body
		}
	}
	for _, rawOrbit := range rawOrbits {
		bodies := strings.Split(rawOrbit, ")")
		parentBody := allBodies[bodies[0]]
		childBody := allBodies[bodies[1]]
		parentBody.AddChild(childBody)
		childBody.AddParent(parentBody)
		allBodies[bodies[1]] = childBody
	}
	return allBodies
}

// CalcShortestJump will give the shortest path.
func CalcShortestJump(givenMap string, body1, body2 string) (jumps int) {
	allbodies := buildMap(givenMap)
	// fmt.Println(allbodies[body1])
	path1 := buildPath(allbodies, allbodies[body1])
	path2 := buildPath(allbodies, allbodies[body2])
	for _, parent1 := range path1 {
		for _, parent2 := range path2 {
			if parent1 == parent2 {
				path1 = path1[1:]
				path2 = path2[1:]
				break
			}
		}
	}

	return len(path1) + len(path2)
}

func buildPath(fullmap map[string]*PlanetaryBody, body *PlanetaryBody) (path []*PlanetaryBody) {
	for body.parent != nil {
		// fmt.Println(body)
		path = append(path, body.parent)
		body = body.parent
	}
	return path
}
