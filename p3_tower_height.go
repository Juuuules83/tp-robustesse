package main

func TowerHeightV1(n int) int {
	var result int

	for i := 1; i <= n; i++ {
		result += i
	}

	return result
}

func TowerHeightV2(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegativeHeight
	}

	if n == 0 {
		return 0, nil
	}

	maxInt := int(^uint(0) >> 1)

	if n == maxInt {
		return 0, ErrOverflow
	}

	a := n
	b := n + 1

	if a%2 == 0 {
		a /= 2
	} else {
		b /= 2
	}

	if a > maxInt/b {
		return 0, ErrOverflow
	}

	return a * b, nil
}
