package main

func TwoSumV1(pile []int, cible int) (int, int, bool) {
	dejavu := make(map[int]int)

	for i, v := range pile {
		if j, ok := dejavu[cible-v]; ok {
			return pile[j], v, true
		}
		dejavu[v] = i
	}

	return 0, 0, false
}
