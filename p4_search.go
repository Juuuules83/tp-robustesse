package main



func SearchV1(ligne []int, v int) (int, error) {
	if len(ligne) == 0 {
		return -1, ErrEmptyPile // Vérification si la pile est vide
	}
	if ligne[0] > ligne[len(ligne)-1] {
		return -1, ErrPileNotSorted
	}
	gauchetab := 0
	droitetab := len(ligne) - 1
	for gauchetab <= droitetab {
		milieutab := (gauchetab + droitetab) / 2
		if ligne[milieutab] == v {
			return milieutab, nil
		}
		if ligne[milieutab] < v {
			gauchetab = milieutab + 1
		} else {
			droitetab = milieutab - 1
		}
	}
	return -1, ErrValueNotFound
}