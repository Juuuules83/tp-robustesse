package main

func DuplicateV1(pile []int) int {

	mapbool := make(map[int]bool)

	for _, nom := range pile {
		if mapbool[nom] {
			return nom
		}
		mapbool[nom] = true
	}
	return 0
}

func DuplicateV2(pile []int) (int, error) {
	n := len(pile) - 1
	if n < 1 {
		return 0, ErrInvalid
	}
	somme := 0
	for _, v := range pile {
		if v < 1 || v > n {
			return 0, ErrInvalid
		}
		somme += v
	}
	return somme - n*(n+1)/2, nil
}
