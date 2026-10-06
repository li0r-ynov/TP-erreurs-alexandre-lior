package kaplas

import (
	"errors"
	"fmt"
	"testing"
)

// Conserver le résultat pour éviter que le compilateur élimine le calcul mesuré.
var sink int

// BenchmarkSmallestV1 mesure la recherche du minimum dans des piles mélangées de tailles croissantes.
func BenchmarkSmallestV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // Préparer les données en dehors de la mesure.
		b.Run(fmt.Sprintf("Shuffled/Size=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := SmallestV1(pile)
				if err != nil {
					b.Fatalf("%s: SmallestV1 returned unexpected error: %v", b.Name(), err)
				}
				sink = got
			}
		})
	}
}

// TestSmallestV1 vérifie que la recherche renvoie le minimum d'une pile non triée.
func TestSmallestUnsortedInput(t *testing.T) {
	t.Run("UnsortedInput", func(t *testing.T) {
		// Préparer une pile dont le minimum se trouve au milieu.
		input := []int{4, 2, 9}
		want := 2

		// Vérifier l'absence d'erreur avant de comparer le minimum obtenu.
		got, err := SmallestV1(input)
		if err != nil {
			t.Fatalf("%s: SmallestV1(%v) returned unexpected error: %v", t.Name(), input, err)
		}
		if got != want {
			t.Errorf("%s: SmallestV1(%v) = %d; want %d", t.Name(), input, got, want)
		}
	})
}

// TestSmallestV1TableDriven vérifie le minimum et l'erreur renvoyés pour chaque cas.
func TestSmallestV1TableDriven(t *testing.T) {
	// Définir les données, le minimum attendu et l'erreur attendue de chaque scénario.
	testCases := []struct {
		name    string
		input   []int
		want    int
		wantErr error
	}{
		{name: "SingleElement", input: []int{7}, want: 7},
		{name: "DuplicateMinimum", input: []int{4, 2, 2, 9}, want: 2},
		{name: "ReverseSortedInput", input: []int{3, 2, 1}, want: 1},
		{name: "NegativeValues", input: []int{-5, 3, -12}, want: -12},
		{name: "EmptyInput", input: []int{}, want: 0, wantErr: ErrEmpty},
		{name: "NilInput", input: nil, want: 0, wantErr: ErrEmpty},
	}

	// Exécuter chaque scénario dans un sous-test pour identifier précisément un échec.
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := SmallestV1(test.input)
			// Comparer les erreurs avec errors.Is, y compris lorsqu'aucune n'est attendue.
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: SmallestV1(%v) error = %v; want %v", t.Name(), test.input, err, test.wantErr)
			}

			// Vérifier aussi la valeur renvoyée pour les piles vides ou nil.
			if got != test.want {
				t.Errorf("%s: SmallestV1(%v) = %d; want %d", t.Name(), test.input, got, test.want)
			}
		})
	}
}
