package main

func QuickSort(scores []int) {
	if len(scores) <= 1 {
		return
	}

	pivot := scores[len(scores)-1]
	i := 0

	for j := 0; j < len(scores)-1; j++ {
		if scores[j] < pivot {
			scores[i], scores[j] = scores[j], scores[i]
			i++
		}
	}

	scores[i], scores[len(scores)-1] = scores[len(scores)-1], scores[i]

	QuickSort(scores[:i])
	QuickSort(scores[i+1:])
}
