package lasagna

func PreparationTime(layers []string, averagePerLayer int) int {
	if averagePerLayer < 1 {
		return len(layers) * 2
	}

	return len(layers) * averagePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, value := range layers {
		switch value {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
	}

	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, numOfPortions int) []float64 {
	newQuantities := make([]float64, len(quantities))

	for i, value := range quantities {
		newQuantities[i] = value * float64(numOfPortions) / 2
	}

	return newQuantities
}
