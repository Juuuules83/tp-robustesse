package main

import (
	"fmt"
	"slices"
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				BubbleSort(scores)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				SelectionSort(scores)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)
		scores := make([]int, n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				InsertionSort(scores)
			}
		})
	}
}

func BenchmarkNearlySorted(b *testing.B) {
	base := NearlySortedScores(100_000)

	b.Run("BubbleSort", func(b *testing.B) {
		scores := make([]int, len(base))

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, base)
			b.StartTimer()

			BubbleSort(scores)
		}
	})

	b.Run("InsertionSort", func(b *testing.B) {
		scores := make([]int, len(base))

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(scores, base)
			b.StartTimer()

			InsertionSort(scores)
		}
	})
}

func BenchmarkMergeSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MergeSort(base)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		random := RandomScores(n)
		sorted := SortedScores(n)

		b.Run(fmt.Sprintf("Random_n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, random)
				b.StartTimer()

				QuickSort(scores)
			}
		})

		b.Run(fmt.Sprintf("Sorted_n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, sorted)
				b.StartTimer()

				QuickSort(scores)
			}
		})
	}
}

func BenchmarkInsertionSortScores(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomPlayers(n)
		players := make([]Score, n)

		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(players, base)
				b.StartTimer()

				InsertionSortScores(players)
			}
		})
	}
}

func BenchmarkPlayerSorts(b *testing.B) {
	base := RandomPlayers(10_000)

	b.Run("InsertionSortScores", func(b *testing.B) {
		players := make([]Score, len(base))

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(players, base)
			b.StartTimer()

			InsertionSortScores(players)
			IsStable(players)
		}
	})

	b.Run("SelectionSortScores", func(b *testing.B) {
		players := make([]Score, len(base))

		for i := 0; i < b.N; i++ {
			b.StopTimer()
			copy(players, base)
			b.StartTimer()

			SelectionSortScores(players)
			IsStable(players)
		}
	})
}

func BenchmarkAllSorts(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		base := RandomScores(n)

		b.Run(fmt.Sprintf("BubbleSort/n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				BubbleSort(scores)
			}
		})

		b.Run(fmt.Sprintf("SelectionSort/n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				SelectionSort(scores)
			}
		})

		b.Run(fmt.Sprintf("InsertionSort/n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				InsertionSort(scores)
			}
		})

		b.Run(fmt.Sprintf("MergeSort/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MergeSort(base)
			}
		})

		b.Run(fmt.Sprintf("QuickSort/n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				QuickSort(scores)
			}
		})

		b.Run(fmt.Sprintf("slices.Sort/n=%d", n), func(b *testing.B) {
			scores := make([]int, n)

			for i := 0; i < b.N; i++ {
				b.StopTimer()
				copy(scores, base)
				b.StartTimer()

				slices.Sort(scores)
			}
		})
	}
}
