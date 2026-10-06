package main

func InsertionSort(scores []int) {
	for i := 1; i < len(scores); i++ {
		key := scores[i]
		j := i - 1

		for j >= 0 && scores[j] > key {
			scores[j+1] = scores[j]
			j--
		}

		scores[j+1] = key
	}
}