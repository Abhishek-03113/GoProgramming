package ScoreRating

func ScoreRating(examScore int) string {
	switch {
	case examScore >= 90:
		return "Excellent"
	case examScore > 75:
		return "Good"
	case examScore > 60:
		return "Pass"
	default:
		return "Fail"

	}
}
