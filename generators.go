package main

import (
	"fmt"
	"math/rand"
)

// Score représente le résultat d'un joueur.
type Score struct {
	Player string
	Score  int
}

var r = rand.New(rand.NewSource(69))

// RandomScores renvoie n scores tirés au hasard entre 0 et 1 000.
func RandomScores(n int) []int {
	scores := make([]int, n)
	for i := range scores {
		scores[i] = r.Intn(1001)
	}
	return scores
}

// SortedScores renvoie n scores tous différents (0, 1, 2…), déjà rangés par ordre croissant.
func SortedScores(n int) []int {
	scores := make([]int, n)
	for i := range scores {
		scores[i] = i
	}
	return scores
}

// ReversedScores renvoie n scores rangés par ordre décroissant (pire cas fréquent).
func ReversedScores(n int) []int {
	scores := SortedScores(n)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		scores[i], scores[j] = scores[j], scores[i]
	}
	return scores
}

// NearlySortedScores renvoie n scores rangés, dont environ 1 % ont été déplacés.
func NearlySortedScores(n int) []int {
	scores := SortedScores(n)
	for k := 0; k < n/100; k++ {
		i, j := r.Intn(n), r.Intn(n)
		scores[i], scores[j] = scores[j], scores[i]
	}
	return scores
}

func RandomPlayers(n int) []Score {
	players := make([]Score, n)
	for i := range players {
		players[i] = Score{Player: fmt.Sprintf("Joueur%05d", i+1), Score: r.Intn(101)}
	}
	return players
}
