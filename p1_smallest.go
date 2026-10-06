package main

func SmallestV1(pile []int) int {
	if len(pile) == 0 {
		return 0
	}

	min := pile[0]
	for i := 1; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]
		}
	}
	return min
}

func SmallestV2(pile []int) int {
	if len(pile) == 0 {
		return 0
	}

	min := pile[0]
	for i := 1; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]
		}
		if min == 1 {
			return 1
		}
	}
	return min
}
