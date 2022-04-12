package test

import "fmt"

func SpiralPattern(n int) string {
	i := 1
	j := 1

	// if n is even
	finalEndI := n/2 + 1
	finalEndJ := n / 2

	// if n is odd
	if n%2 == 1 {
		finalEndJ = n/2 + 1
	}
	upperFence := n
	lowerFence := 1

	// direction
	dir := 'r'
	res := ""

	for i != finalEndI || j != finalEndJ {
		switch dir {
		case 'r':
			for j <= upperFence {
				res += fmt.Sprintf("%d ", n*(i-1)+j)

				j++
			}
			j--
			i++
			dir = 'd'
		case 'd':
			for i <= upperFence {
				res += fmt.Sprintf("%d ", n*(i-1)+j)
				i++
			}
			i--
			j--
			dir = 'l'
		case 'l':
			for j >= lowerFence {
				res += fmt.Sprintf("%d ", n*(i-1)+j)
				j--
			}
			j++
			i--
			dir = 'u'
		case 'u':
			lowerFence++
			upperFence--
			for i >= lowerFence {
				res += fmt.Sprintf("%d ", n*(i-1)+j)
				i--
			}
			i++
			j++
			dir = 'r'
		}
	}

	// print the final element in spiral pattern
	res += fmt.Sprint((i-1)*n + j)

	return res
}
