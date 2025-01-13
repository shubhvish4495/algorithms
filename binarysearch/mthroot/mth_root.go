package main

import (
	"fmt"
	"math"
)

func nthRoot(n, m int) int {
	l := 1
	h := m

	for l <= h {
		mid := (l + h) / 2
		pVal := int(math.Pow(float64(mid), float64(n)))
		if pVal == m {
			return mid
		}

		if pVal > m {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func main() {
	n := 3
	m := 27

	fmt.Println(nthRoot(n, m))
}
