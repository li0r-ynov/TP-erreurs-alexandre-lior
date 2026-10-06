package main

func CountV1(pile []int, plafond int) ([]int, error) {
    if len(pile) == 0 {
		return nil, ErrEmptyPile
	}
    slice := make([]int, plafond+1)
    for _, nom := range pile  {
        if nom <= plafond {
            slice[nom]++}
    }
    return slice, nil
}