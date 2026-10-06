package main

import (
	"errors"
	"reflect"
	"testing"
)

type CasesTable struct {
	name    string
	input   []int
	value   int
	want    int
	wantErr error
}

type Ntable struct {
	name    string
	input   int
	want    int
	wantErr error
}

type CountCase struct {
	name    string
	input   []int
	value   int
	want    []int
	wantErr error
}

type FirstUniqueCase struct {
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
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: SmallestV1(%v) = %d; want %d", t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestDuplicateTableDriven(t *testing.T) {
	testCases := []CasesTable{
		{name: "SingleDuplicate", input: []int{7, 7}, want: 7},
		{name: "DuplicateMinimum", input: []int{4, 2, 2, 9}, want: 2},
		{name: "DuplicateAtEnd", input: []int{3, 2, 1, 1}, want: 1},
		{name: "DuplicateNegative", input: []int{-5, 3, -12, -12}, want: -12},
		{name: "EmptyInput", input: []int{}, want: -1, wantErr: ErrEmptyPile},
		{name: "NilInput", input: nil, want: -1, wantErr: ErrEmptyPile},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := DuplicateV1(test.input)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: DuplicateV1(%v) error = %v; want %v",
					t.Name(), test.input, err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: DuplicateV1(%v) = %d; want %d",
					t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestTowerHeightTableDriven(t *testing.T) {
	testCases := []Ntable{
		{name: "SingleFloor", input: 1, want: 1},
		{name: "TwoFloors", input: 2, want: 3},
		{name: "FiveFloors", input: 5, want: 15},
		{name: "TenFloors", input: 10, want: 55},
		{name: "ZeroFloors", input: 0, want: -1, wantErr: ErrEmptyPile},
		{name: "NegativeFloors", input: -1, want: 0},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := TowerHeightV1(test.input)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: TowerHeightV1(%v) error = %v; want %v",
					t.Name(), test.input, err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: TowerheightV1(%v) = %d; want %d",
					t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestSearchTableDriven(t *testing.T) {
	testCases := []CasesTable{
		{name: "FindFirst", input: []int{1, 2, 3, 4, 5}, value: 1, want: 0},
		{name: "FindMiddle", input: []int{1, 2, 3, 4, 5}, value: 3, want: 2},
		{name: "FindLast", input: []int{1, 2, 3, 4, 5}, value: 5, want: 4},
		{name: "NotFound", input: []int{1, 2, 3, 4, 5}, value: 9, want: -1},
		{name: "EmptyInput", input: []int{}, value: 1, want: -1, wantErr: ErrEmptyPile},
		{name: "NilInput", input: nil, value: 1, want: -1, wantErr: ErrEmptyPile},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := SearchV1(test.input, test.value)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: SearchV1(%v) error = %v; want %v",
					t.Name(), test.input, err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: SearchV1(%v) = %d; want %d",
					t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestCountTableDriven(t *testing.T) {
	testCases := []CountCase{
		{"SimpleCount", []int{1, 2, 2, 4}, 4, []int{0, 1, 2, 0, 1}, nil},
		{"AllSame", []int{2, 2, 2}, 2, []int{0, 0, 3}, nil},
		{"NoDuplicate", []int{1, 2, 3}, 3, []int{0, 1, 1, 1}, nil},
		{"EmptyInput", []int{}, 3, nil, ErrEmptyPile},
		{"NilInput", nil, 3, nil, ErrEmptyPile},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := CountV1(test.input, test.value)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: CountV1(%v) error = %v; want %v",
					t.Name(), test.input, err, test.wantErr)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: CountV1(%v) = %d; want %d",
					t.Name(), test.input, got, test.want)
			}
		})
	}
}

func TestFirstUniqueTableDriven(t *testing.T) {
	testCases := []FirstUniqueCase{
		{"SimpleUnique", []int{1, 2, 2, 4}, 1, nil},
		{"UniqueAtEnd", []int{2, 2, 4}, 4, nil},
		{"AllSame", []int{2, 2, 2}, -1, nil},
		{"NoUnique", []int{2, 2, 3, 3}, -1, nil},
		{"NoDuplicate", []int{1, 2, 3}, 1, nil},
		{"EmptyInput", []int{}, -1, ErrEmptyPile},
		{"NilInput", nil, -1, ErrEmptyPile},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := FirstUniqueV1(test.input)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf("%s: FirstUniqueV1(%v) error = %v; want %v",
					t.Name(), test.input, err, test.wantErr)
			}

			if got != test.want {
				t.Errorf("%s: FirstUniqueV1(%v) = %d; want %d",
					t.Name(), test.input, got, test.want)
			}
		})
	}
}
