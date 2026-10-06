package main

func FirstUniqueV1(ligne []int) int {

	compteur := make(map[int]int)

	for _, v := range ligne {
		compteur[v]++
	}

	for _, v := range ligne {
		if compteur[v] == 1 {
			return v
		}

	}
	return -1

}
