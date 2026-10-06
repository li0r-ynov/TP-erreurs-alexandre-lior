package main

func DuplicateV1(pile []int) int {
	double := 0
	for i := 0 ; i<len(pile);i++{
		for j := 1 ; j<len(pile);j++{
			if pile[i] == pile[j] && i != j{
				double = pile[i]
				return double
			}
		}
	}
	return -1
}

func DuplicateV2 (pile []int) int {
	maps := make(map[int]int)
	result := 0
	for i := 0 ; i<len(pile);i++{
		maps[pile[i]]+=1
	if maps[pile[i]] == 2 {
		result = pile[i]
		break
	}
	}
	return result
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
