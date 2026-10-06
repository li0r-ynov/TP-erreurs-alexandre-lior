package main

func MergeSort(scores []int) []int {
	if len(scores) <= 1 {
		return scores
	}

	mid := len(scores) / 2

	left := MergeSort(scores[:mid])
	right := MergeSort(scores[mid:])

	result := make([]int, 0, len(scores))

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}