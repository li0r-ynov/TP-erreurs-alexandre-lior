package kaplas

import "math/rand"

var r = rand.New(rand.NewSource(42))

// Sorted renvoie les Kaplas numérotés de 1 à n, dans l'ordre.
func Sorted(n int) []int {
	pile := make([]int, n)
	for i := range pile {
		pile[i] = i + 1
	}
	return pile
}

// Shuffled renvoie les Kaplas numérotés de 1 à n, mélangés.
func Shuffled(n int) []int {
	pile := Sorted(n)
	r.Shuffle(n, func(i, j int) { pile[i], pile[j] = pile[j], pile[i] })
	return pile
}

// WithDuplicate renvoie les Kaplas de 1 à n mélangés,
// puis une copie du numéro d placée en dernier (cas défavorable).
func WithDuplicate(n, d int) []int {
	return append(Shuffled(n), d)
}

// Random renvoie n Kaplas aux numéros tirés au hasard entre 1 et plafond.
func Random(n, plafond int) []int {
	pile := make([]int, n)
	for i := range pile {
		pile[i] = r.Intn(plafond) + 1
	}
	return pile
}

// WithTwins renvoie environ n Kaplas où chaque numéro a un jumeau,
// sauf le dernier Kapla, seul de son numéro (cas défavorable).
func WithTwins(n int) []int {
	moitie := (n - 1) / 2
	pile := append(Shuffled(moitie), Shuffled(moitie)...)
	r.Shuffle(len(pile), func(i, j int) {
		pile[i], pile[j] = pile[j], pile[i]
	})
	return append(pile, moitie+1)
}
