package main

// GetModuleFuelRequirement Calculates how much fuel is required on a
// module basis.
func GetModuleFuelRequirement(mass int) int {
	if mass < 9 {
		return 0
	}
	fuelRequired := (mass / 3) - 2
	fuelRequired = fuelRequired + GetModuleFuelRequirement(fuelRequired)
	return fuelRequired
}

// GetFuelTotal will calculate the total fuel required to launch.
func GetFuelTotal(modulemasses []int) int {
	totalFuel := 0
	for _, mass := range modulemasses {
		totalFuel = totalFuel + GetModuleFuelRequirement(mass)
	}
	return totalFuel
}
