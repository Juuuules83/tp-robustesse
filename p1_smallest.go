package main

func SmallestV2(pile []int) (int, error) {
	if len(pile) == 0 {
		return 0, ErrEmptyPile // Vérification si la pile est vide
	}

	min := pile[0]

	for i := 1; i < len(pile); i++ {
		if pile[i] < min {
			min = pile[i]
		}
	}

	return min, nil
}