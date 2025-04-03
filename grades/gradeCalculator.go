package grades

// CalculateAverage computes the average of given grades safely.
func CalculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0.0
	}

	var sum float64
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}
