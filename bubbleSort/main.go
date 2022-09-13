// bubblesort implementation

package main

import "fmt"

func sort(x []int) []int {
	ix := x
	for i := range ix {
		for ii := range ix {
			// first loop <; then will be ordered asc
			// first loop >; then will be ordered desc
			if ix[i] < ix[ii] {
				ix[i], ix[ii] = ix[ii], ix[i]
			}
		}
	}
	return ix
}

func main() {
	a := []int{8, 5, 7, 1, 9, 6, 5, 4, 3, 2}

	a = sort(a)

	fmt.Println(a)
}
