package lasagna

func PreparationTime(layers []string, time int) int {
	var numOfLayers = len(layers)
	if time <= 0 {
		return numOfLayers * 2
	}
	return numOfLayers * time
}

func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendLayers, layers []string) {
	layers[len(layers)-1] = friendLayers[len(friendLayers)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	size := len(quantities)
	multiplier := float64(portions) * 0.5
	scaledQuantities := make([]float64, size)
	for i := 0; i < size; i++ {
		scaledQuantities[i] = quantities[i] * multiplier
	}
	return scaledQuantities
}
