package main


func CountV1(pile []int, plafond int) ([]int, error) {
	if plafond < 0 {
		return nil, ErrInvalid
	}
	slice := make([]int, plafond+1)
	for _, v := range pile {
		if v < 0 {
			return nil, ErrInvalid
		}
		if v <= plafond {
			slice[v]++
		}
	}
	return slice, nil
}