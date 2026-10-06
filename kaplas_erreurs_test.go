package main

import (
	"errors"
	"testing"
)

type CasesTable struct {
	name    string
	input   []int
	want    int
	wantErr error
}

func TestSmallestTableDriven(t *testing.T) {
	// Définir les données, le minimum attendu et l'erreur attendue de chaque scénario.
	testCases := []CasesTable{
		{name: "SingleElement", input: []int{7}, want: 7},
		{name: "DuplicateMinimum", input: []int{4, 2, 2, 9}, want: 2},
		{name: "ReverseSortedInput", input: []int{3, 2, 1}, want: 1},
		{name: "NegativeValues", input: []int{-5, 3, -12}, want: -12},
		{name: "EmptyInput", input: []int{}, want: -1, wantErr: ErrEmptyPile},
		{name: "NilInput", input: nil, want: -1, wantErr: ErrEmptyPile},
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

func TestXXXXXTableDriven(t *testing.T) {
	// Définir les données, le minimum attendu et l'erreur attendue de chaque scénario.
	testCases := []CasesTable{
		{name: "SingleElement", input: []int{7}, want: 7},
		{name: "DuplicateMinimum", input: []int{4, 2, 2, 9}, want: 2},
		{name: "ReverseSortedInput", input: []int{3, 2, 1}, want: 1},
		{name: "NegativeValues", input: []int{-5, 3, -12}, want: -12},
		{name: "EmptyInput", input: []int{}, want: -1, wantErr: ErrEmptyPile},
		{name: "NilInput", input: nil, want: -1, wantErr: ErrEmptyPile},
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
