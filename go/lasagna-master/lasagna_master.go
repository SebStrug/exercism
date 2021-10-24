package lasagna

// Preparation Time returns the total preparation time for all layers
// accepts the layers and the average preparation time per layer in minutes
func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// Quantities returns the number of noodles and sauce required given the layers
// Each noodle layer requires 50g noodles. Each sauce layer requires 0.2l of sauce
func Quantities(layers []string) (int, float64) {
	var totalNoodles int
	var sauce float64
	for _, l := range layers {
		switch {
		case l == "noodles":
			totalNoodles += 50
		case l == "sauce":
			sauce += 0.2
		}
	}
	return totalNoodles, sauce
}

// AddSecretIngredient adds the last item of the second array to the first array
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// ScaleRecipe finds the amounts required to cook lasagna for multiple people
// amount is the quantity required to feed 2, portions is the number of people to feed
func ScaleRecipe(amounts []float64, portions int) []float64 {
	multiplier := float64(portions) / 2
	scaledAmounts := make([]float64, len(amounts))
	for ind, item := range amounts {
		scaledAmounts[ind] = item * float64(multiplier)
	}
	return scaledAmounts
}
