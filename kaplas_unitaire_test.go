package main

import (
	"errors"
	"testing"
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
