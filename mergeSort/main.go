// merge sort implementation
// time complexity is O(n*Log n)
package main

import "fmt"

func mergeSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	first := mergeSort(x[:len(x)/2])
	second := mergeSort(x[len(x)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	// check if there is any left over and add it to the end
	if i < len(a) {
		final = append(final, a[i:]...)
	}
	if j < len(b) {
		final = append(final, b[j:]...)
	}

	return final
}

func main() {
	a := []int{5, 7, 2, 1, 6, 3, 10, 8, 41, 20, 1, 0, 5, 18, 21}
	b := mergeSort(a)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(b)
}
