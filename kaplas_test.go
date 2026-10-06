package main

import (
	"fmt"
	"testing"
)

var sink int // garde le résultat pour que le compilateur ne supprime pas l'appel

func BenchmarkSmallestV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, _ = SmallestV2(pile)
			}
		})
	}
}

func BenchmarkDuplicateV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := WithDuplicate(n, 3) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, _ = DuplicateV2(pile)
			}
		})
	}
}

func BenchmarkTowerHeightV2(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		b.Run(fmt.Sprintf("V2/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, _ = TowerHeightV2(n)
			}
		})
	}
}

func BenchmarkCountV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Random(n, n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				res, _ := CountV1(pile, n)
				sink = len(res)
			}
		})
	}
}

func BenchmarkSearchV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := Sorted(n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, _ = SearchV1(ligne, n+1)
			}
		})
	}
}

func BenchmarkFirstUniqueV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := WithTwins(n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink, _ = FirstUniqueV1(ligne)
			}
		})
	}
}

func BenchmarkTwoSumV1(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		ligne := Shuffled(n)
		cible := 3 * n
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				resultat, _, _ := TwoSumV1(ligne, cible)
				sink = resultat
			}
		})
	}
}
