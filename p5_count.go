package main

func CountV1(pile []int, plafond int) []int {
    slice := make([]int, plafond+1)
    for _, nom := range pile  {
        if nom <= plafond {
            slice[nom]++}
    }
    return slice
}