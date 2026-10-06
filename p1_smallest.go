package main

func SmallestV1(pile []int) (int, error) {
	if len(pile) == 0 {
		return -1, ErrEmptyPile
	}
	min := pile[0]
	for i := 1; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]
		}
	}
	return min, nil
}
