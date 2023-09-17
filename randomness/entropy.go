package randomness

import "math"

func CalculateNonUniformEntropy(args ...float64) float64 {
	entropy := 0.0
	for i := 0; i < len(args); i++ {
		entropy += -args[i] * math.Log2(args[i])
	}
	return entropy
}

func CalculateUniformEntropy(probability float64, occurrence int) float64 {
	entropy := 0.0
	entropy += -probability * math.Log2(probability)
	entropy *= float64(occurrence)
	return entropy
}
