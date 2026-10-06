package main

func SelectionSortScores(players []Score) {
	for i := 0; i < len(players)-1; i++ {
		maxIndex := i

		for j := i + 1; j < len(players); j++ {
			if players[j].Score > players[maxIndex].Score {
				maxIndex = j
			}
		}

		players[i], players[maxIndex] = players[maxIndex], players[i]
	}
}

func IsStable(players []Score) bool {
	for i := 0; i < len(players)-1; i++ {
		if players[i].Score == players[i+1].Score {
			if players[i].Player > players[i+1].Player {
				return false
			}
		}
	}

	return true
}
