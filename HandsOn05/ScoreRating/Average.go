package ScoreRating

func Average(nums ...float64) float64 {
	sum := 0.0
	cnt := 1.0
	for _, vals := range nums {
		sum += vals
		cnt++
	}

	return sum / cnt
}
