package main

func FirstUniqueV1(ligne []int) (int, error) {
	if len(ligne) == 0 {
    return -1, ErrEmptyPile
}

	compteur := make(map[int]int)

	for _, v := range ligne {
		compteur[v]++
	}

	for _, v := range ligne {
		if compteur[v] == 1 {
			return v, nil
		}

	}
	return -1,nil

}
