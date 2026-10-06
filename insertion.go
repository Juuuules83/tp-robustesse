package main

func InsertionSort(scores []int) []int {
	for i := 1; i < len(scores); i++ {
		key := scores[i]
		j := i - 1
		for j >= 0 && scores[j] > key {
			scores[j+1] = scores[j]
			j--
		}
		scores[j+1] = key
	}
	return scores
}

func InsertionSortScores(players []Score) {
	for i := 1; i < len(players); i++ {
		key := players[i]
		j := i - 1
		for j >= 0 && players[j].Score > key.Score {
			players[j+1] = players[j]
			j--
		}
		players[j+1] = key
	}
}
