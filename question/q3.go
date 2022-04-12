package test

func ProducePiOf(a []int) []int {
	res := make([]int, len(a))
	proOfAll := 1
	noZero := true
	indexFirstZero := -1

	for i, v := range a {
		if v == 0 {
			if noZero {
				indexFirstZero = i
				noZero = false
			} else {
				return res
			}
		} else {
			proOfAll *= v
		}
	}

	if noZero {
		for i := 0; i < len(a); i++ {
			res[i] = proOfAll / a[i]
		}
	} else {
		res[indexFirstZero] = proOfAll
	}

	return res
}
