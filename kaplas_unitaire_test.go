package main

import (
	"errors"
	"testing"
	"reflect"
)


// ------ SMALLESTV2 TESTS ------ //
func TestSmallestV2(t *testing.T) {
	testCases := []struct {
		name    string
		pile    []int
		want    int
		wantErr error
	}{
		{
			name: "Normal",
			pile: []int{4, 2, 9},
			want: 2,
		},
		{
			name: "SingleElement",
			pile: []int{7},
			want: 7,
		},
		{
			name: "NegativeValues",
			pile: []int{-5, 3, -12},
			want: -12,
		},
		{
			name: "DuplicateMinimum",
			pile: []int{4, 2, 2, 9},
			want: 2,
		},
		{
			name:    "EmptyInput",
			pile:    []int{},
			want:    0,
			wantErr: ErrEmptyPile,
		},
		{
			name: "MinimumIsOne",
			pile: []int{4, 3, 1, 8},
			want: 1,
		},
		{
			name: "OneThenNegative",
			pile: []int{4, 3, 1, -5},
			want: -5,
		},
		{
			name:    "NilInput",
			pile:    nil,
			want:    0,
			wantErr: ErrEmptyPile,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := SmallestV2(test.pile)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"SmallestV2(%v) error = %v, want %v",
					test.pile,
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"SmallestV2(%v) = %d, want %d",
					test.pile,
					got,
					test.want,
				)
			}
		})
	}
}



// ------ DUPLICATEV2 TESTS ------ //
func TestDuplicateV2(t *testing.T) {
	testCases := []struct {
		name    string
		pile    []int
		want    int
		wantErr error
	}{
		{
			name: "Normal",
			pile: []int{1, 2, 2, 3},
			want: 2,
		},
		{
			name: "DuplicateAtBeginning",
			pile: []int{1, 1, 2, 3},
			want: 1,
		},
		{
			name: "DuplicateAtEnd",
			pile: []int{1, 2, 3, 3},
			want: 3,
		},
		{
			name: "LargeValues",
			pile: []int{1, 2, 3, 4, 5, 5},
			want: 5,
		},
		{
			name:    "ZeroValue",
			pile:    []int{0, 2, 2},
			want:    0,
			wantErr: ErrInvalid,
		},
		{
			name:    "NegativeValue",
			pile:    []int{-1, 2, 2},
			want:    0,
			wantErr: ErrInvalid,
		},
		{
			name:    "ValueTooLarge",
			pile:    []int{1, 2, 4},
			want:    0,
			wantErr: ErrInvalid,
		},
		{
			name:    "EmptyInput",
			pile:    []int{},
			want:    0,
			wantErr: ErrInvalid,
		},
		{
			name:    "SingleElement",
			pile:    []int{1},
			want:    0,
			wantErr: ErrInvalid,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := DuplicateV2(test.pile)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"DuplicateV2(%v) error = %v, want %v",
					test.pile,
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"DuplicateV2(%v) = %d, want %d",
					test.pile,
					got,
					test.want,
				)
			}
		})
	}
}





// ------ TOWERHEIGHTV2 TESTS ------ //
func TestTowerHeightV2(t *testing.T) {
	testCases := []struct {
		name    string
		n       int
		want    int
		wantErr error
	}{
		{
			name: "Normal",
			n:    5,
			want: 15,
		},
		{
			name: "Zero",
			n:    0,
			want: 0,
		},
		{
			name: "One",
			n:    1,
			want: 1,
		},
		{
			name: "LargeValue",
			n:    100,
			want: 5050,
		},
		{
			name:    "NegativeValue",
			n:       -1,
			want:    0,
			wantErr: ErrNegativeHeight,
		},
		{
			name:    "Overflow",
			n:       int(^uint(0)>>1),
			want:    0,
			wantErr: ErrOverflow,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := TowerHeightV2(test.n)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"TowerHeightV2(%d) error = %v, want %v",
					test.n,
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"TowerHeightV2(%d) = %d, want %d",
					test.n,
					got,
					test.want,
				)
			}
		})
	}
}





// ------ SEARCHV1 TESTS ------ //
func TestSearchV1(t *testing.T) {
	testCases := []struct {
		name    string
		ligne   []int
		v       int
		want    int
		wantErr error
	}{
		{
			name:  "ValueFound",
			ligne: []int{-4, -2, 0, 1, 3, 5, 7},
			v:     3,
			want:  4,
		},
		{
			name:  "FirstValue",
			ligne: []int{-4, -2, 0, 1, 3, 5, 7},
			v:     -4,
			want:  0,
		},
		{
			name:  "LastValue",
			ligne: []int{-4, -2, 0, 1, 3, 5, 7},
			v:     7,
			want:  6,
		},
		{
			name:    "EmptyInput",
			ligne:   []int{},
			v:       3,
			want:    -1,
			wantErr: ErrEmptyPile,
		},
		{
			name:    "NotSorted",
			ligne:   []int{7, 3, 5, 1},
			v:       3,
			want:    -1,
			wantErr: ErrPileNotSorted,
		},
		{
			name:    "ValueNotFound",
			ligne:   []int{-4, -2, 0, 1, 3, 5, 7},
			v:       10,
			want:    -1,
			wantErr: ErrValueNotFound,
		},
		{
			name:    "NilInput",
			ligne:   nil,
			v:       3,
			want:    -1,
			wantErr: ErrEmptyPile,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := SearchV1(test.ligne, test.v)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"SearchV1(%v, %d) error = %v, want %v",
					test.ligne,
					test.v,
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"SearchV1(%v, %d) = %d, want %d",
					test.ligne,
					test.v,
					got,
					test.want,
				)
			}
		})
	}
}


// ------ FIRSTUNIQUEV1 TESTS ------ //
func TestFirstUniqueV1(t *testing.T) {
	testCases := []struct {
		name    string
		ligne   []int
		want    int
		wantErr error
	}{
		{
			name:  "Normal",
			ligne: []int{4, 2, 4, 3, 2},
			want:  3,
		},
		{
			name:  "FirstElementUnique",
			ligne: []int{1, 2, 2, 3, 3},
			want:  1,
		},
		{
			name:  "LastElementUnique",
			ligne: []int{1, 1, 2, 2, 3},
			want:  3,
		},
		{
			name:  "NegativeValues",
			ligne: []int{-1, 2, -1, 3, 2},
			want:  3,
		},
		{
			name:  "SingleElement",
			ligne: []int{7},
			want:  7,
		},
		{
			name:    "EmptyInput",
			ligne:   []int{},
			want:    0,
			wantErr: ErrEmptyPile,
		},
		{
			name:    "NilInput",
			ligne:   nil,
			want:    0,
			wantErr: ErrEmptyPile,
		},
		{
			name:    "NoUniqueValue",
			ligne:   []int{1, 1, 2, 2, 3, 3},
			want:    0,
			wantErr: ErrValueNotFound,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := FirstUniqueV1(test.ligne)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"FirstUniqueV1(%v) error = %v, want %v",
					test.ligne,
					err,
					test.wantErr,
				)
			}

			if got != test.want {
				t.Errorf(
					"FirstUniqueV1(%v) = %d, want %d",
					test.ligne,
					got,
					test.want,
				)
			}
		})
	}
}

// ------ COUNTV1 TESTS ------ //
func TestCountV1(t *testing.T) {
	testCases := []struct {
		name    string
		pile    []int
		plafond int
		want    []int
		wantErr error
	}{
		{
			name:    "Normal",
			pile:    []int{0, 1, 1, 2, 2, 2},
			plafond: 2,
			want:    []int{1, 2, 3},
		},
		{
			name:    "ZeroPlafond",
			pile:    []int{0, 0, 0},
			plafond: 0,
			want:    []int{3},
		},
		{
			name:    "ValuesAbovePlafond",
			pile:    []int{0, 1, 2, 5, 7},
			plafond: 2,
			want:    []int{1, 1, 1},
		},
		{
			name:    "EmptyInput",
			pile:    []int{},
			plafond: 3,
			want:    []int{0, 0, 0, 0},
		},
		{
			name:    "NilInput",
			pile:    nil,
			plafond: 3,
			want:    []int{0, 0, 0, 0},
		},
		{
			name:    "NegativeValue",
			pile:    []int{1, 2, -1, 3},
			plafond: 3,
			want:    nil,
			wantErr: ErrInvalid,
		},
		{
			name:    "NegativePlafond",
			pile:    []int{1, 2, 3},
			plafond: -1,
			want:    nil,
			wantErr: ErrInvalid,
		},
		{
			name:    "DuplicateValues",
			pile:    []int{1, 1, 1, 2, 2},
			plafond: 2,
			want:    []int{0, 3, 2},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := CountV1(test.pile, test.plafond)

			if !errors.Is(err, test.wantErr) {
				t.Fatalf(
					"CountV1(%v, %d) error = %v, want %v",
					test.pile,
					test.plafond,
					err,
					test.wantErr,
				)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(
					"CountV1(%v, %d) = %v, want %v",
					test.pile,
					test.plafond,
					got,
					test.want,
				)
			}
		})
	}
}

// ------ TWOSUMV1 TESTS ------ //
func TestTwoSumV1(t *testing.T) {
	testCases := []struct {
		name      string
		pile      []int
		cible     int
		wantA     int
		wantB     int
		wantFound bool
	}{
		{
			name:      "Normal",
			pile:      []int{2, 7, 11, 15},
			cible:     9,
			wantA:     2,
			wantB:     7,
			wantFound: true,
		},
		{
			name:      "PairAtEnd",
			pile:      []int{1, 4, 6, 10},
			cible:     16,
			wantA:     6,
			wantB:     10,
			wantFound: true,
		},
		{
			name:      "NegativeValues",
			pile:      []int{-3, 5, 8, 2},
			cible:     5,
			wantA:     -3,
			wantB:     8,
			wantFound: true,
		},
		{
			name:      "DuplicateValues",
			pile:      []int{3, 3, 7},
			cible:     6,
			wantA:     3,
			wantB:     3,
			wantFound: true,
		},
		{
			name:      "NoPair",
			pile:      []int{1, 2, 3, 4},
			cible:     20,
			wantA:     0,
			wantB:     0,
			wantFound: false,
		},
		{
			name:      "EmptyInput",
			pile:      []int{},
			cible:     5,
			wantA:     0,
			wantB:     0,
			wantFound: false,
		},
		{
			name:      "SingleElement",
			pile:      []int{5},
			cible:     5,
			wantA:     0,
			wantB:     0,
			wantFound: false,
		},
		{
			name:      "SameValueDifferentPositions",
			pile:      []int{1, 2, 2, 4},
			cible:     4,
			wantA:     2,
			wantB:     2,
			wantFound: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			gotA, gotB, gotFound := TwoSumV1(test.pile, test.cible)

			if gotA != test.wantA {
				t.Errorf(
					"TwoSumV1(%v, %d) first value = %d, want %d",
					test.pile,
					test.cible,
					gotA,
					test.wantA,
				)
			}

			if gotB != test.wantB {
				t.Errorf(
					"TwoSumV1(%v, %d) second value = %d, want %d",
					test.pile,
					test.cible,
					gotB,
					test.wantB,
				)
			}

			if gotFound != test.wantFound {
				t.Errorf(
					"TwoSumV1(%v, %d) found = %t, want %t",
					test.pile,
					test.cible,
					gotFound,
					test.wantFound,
				)
			}
		})
	}
}