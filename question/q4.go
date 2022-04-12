package test

func PaintMatrixWith1(a [][]int) [][]int {
	r := make([]bool, len(a))
	c := make([]bool, len(a[0]))

	// mark rows and cols which have 1
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] == 1 {
				r[i] = true
				c[j] = true
			}
		}
	}

	// paint row
	for i := 0; i < len(a); i++ {
		if r[i] {
			for j := 0; j < len(a[0]); j++ {
				a[i][j] = 1
			}
		}
	}

	// paint col
	for j := 0; j < len(a[0]); j++ {
		if c[j] {
			for i := 0; i < len(a); i++ {
				a[i][j] = 1
			}
		}
	}

	return a
}
