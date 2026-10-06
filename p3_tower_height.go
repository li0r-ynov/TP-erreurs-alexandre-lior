package main

func TowerHeightV1(n int) (int, error) {
	if n == 0 {
		return -1, ErrEmptyPile
	}
	var result int

	for i := 1; i <= n; i++ {
		result += i
	}

	return result, nil
}
