package main

func BubbleSort(scores []int) []int {
	for lastIndex := len(scores) - 1; lastIndex > 0; lastIndex-- {
		swapped := false
		for i := 0; i < lastIndex; i++ {
			if scores[i] > scores[i+1] {
				scores[i], scores[i+1] = scores[i+1], scores[i]
				swapped = true
			}
		}
		if !swapped {
			return scores
		}
	}
	return scores
}
