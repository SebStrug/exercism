package lasagna

const OvenTime int = 40

func RemainingOvenTime(minsInOven int) int {
	return OvenTime - minsInOven
}

func PreparationTime(numLayers int) int {
	return numLayers * 2
}

// TODO: define the 'ElapsedTime()' function
func ElapsedTime(numLayers int, minsInOven int) int {
	return PreparationTime(numLayers) + minsInOven
}
