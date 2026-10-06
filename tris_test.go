package main

import (
	"fmt"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := NearlySortedScores(n) // préparation hors de la mesure
		scores := make([]int, n)
		b.Run(fmt.Sprintf("/BubbleSort/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()
				BubbleSort(scores)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		scores := RandomScores(n)
		b.Run(fmt.Sprintf("/MergeSort/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MergeSort(scores)
			}
		})
	}
}
