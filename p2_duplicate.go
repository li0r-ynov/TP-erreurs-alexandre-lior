package main

func DuplicateV1(pile []int) (int, error) {
    if len(pile) == 0 {
        return -1, ErrEmptyPile
    }

    for i := 0; i < len(pile); i++ {
        for j := i + 1; j < len(pile); j++ {
            if pile[i] == pile[j] {
                return pile[i], nil
            }
        }
    }

    return -1, nil
}


/* package main

func DuplicateV1(pile []int) (int,error) {

	mapbool := make(map[int]bool)

	for _, nom := range pile {
		if mapbool[nom] {
			return nom,nil
		}
		mapbool[nom] = true
	}
	return 0,nil
}

func DuplicateV2(pile []int) int {
	n := len(pile) - 1

	sommeT := (n * (n + 1)) / 2

	sommeR := 0

	for _, nom := range pile {
		sommeR += nom
	}
	return sommeR - sommeT
} */
