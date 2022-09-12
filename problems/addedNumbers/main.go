// given an unsorted array of numbers, find 2 elements that sums to k

package main

import "fmt"

func doesSumExist(x []int, val int) bool {
	found := make([]int, len(x))
	for _, v := range x {
		if found[v] != 0 {
			return true
		}
		found = append(found, val-v)
	}
	return false
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 10

	// brute force O(n^2)
	// for i, v := range input {
	// 	for j := i + 1; j < len(input); j++ {
	// 		if v+input[j] == k {
	// 			fmt.Printf("%v + %v = %v\n", v, input[j], k)
	// 			return
	// 		}
	// 	}
	// }

	// hashmap O(n)
	// remainder := make(map[int]int)
	// for _, v := range input {
	// 	if _, ok := remainder[k-v]; ok {
	// 		fmt.Printf("%v + %v = %v\n", k-v, v, k)
	// 		return
	// 	}
	// 	remainder[v] = k - v
	// }

	// fmt.Printf("No numbers add up to %v", k)

	fmt.Println(doesSumExist(input, k))

}
