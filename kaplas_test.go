package main

import (
	"fmt"
	"testing"
)

var sink int // garde le résultat pour que le compilateur ne supprime pas l'appel
func BenchmarkSmallestV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SmallestV1(pile)
			}
		})
	}
}

func BenchmarkSmallestV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SmallestV2(pile)
			}
		})
	}
}
func BenchmarkDuplicateV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := WithDuplicate(n, 3) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = DuplicateV1(pile)
			}
		})
	}
}
func BenchmarkDuplicateV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := WithDuplicate(n, 3) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = DuplicateV2(pile)
			}
		})
	}
}
func BenchmarkTowerHeightV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = TowerHeightV1(n)
			}
		})
	}
}
func BenchmarkTowerHeightV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = TowerHeightV2(n)
			}
		})
	}
}

func BenchmarkCountV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Random(n, n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = len(CountV1(pile, n))
			}
		})
	}
}

func BenchmarkSearchV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := Sorted(n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SearchV1(ligne, n+1)
			}
		})
	}
}

func BenchmarkFirstUniqueV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := WithTwins(n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = FirstUniqueV1(ligne)
			}
		})
	}
}

func BenchmarkTwoSumV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := Shuffled(n)
		cible := 3*n
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultat, _, _ := TwoSumV1(ligne, cible)
				sink = resultat
			}
		})
	}
}
