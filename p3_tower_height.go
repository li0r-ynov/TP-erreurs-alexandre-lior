package main

func TowerHeightV1(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

func TowerHeightV2(n int) int {
	return n * (n + 1) / 2
}
