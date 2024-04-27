package main

import "fmt"

func main() {
	a := make([][]int, 0)
	numRows := 5

	if numRows >= 1 {
		a = append(a, []int{1})
	}

	if numRows >= 2 {
		a = append(a, []int{1, 1})
	}

	for i := 3; i <= numRows; i++ {
		tempArr := make([]int, 0)
		tempArr = append(tempArr, 1)
		for j := range len(a[i-2]) - 1 {
			tempArr = append(tempArr, a[i-2][j]+a[i-2][j+1])
		}

		tempArr = append(tempArr, 1)

		a = append(a, tempArr)
	}

	fmt.Println(a)
}
