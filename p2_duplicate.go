package main

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
}
