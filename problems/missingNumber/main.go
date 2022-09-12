// find the missing number in the array from 1 to n
// Questions:
// Will there always only be one missing number? Will there always be a missing number?

package main

import "fmt"

func sum(x []int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func makeRange(min, max int) []int {
	var xRange []int
	for i := min; i <= max; i++ {
		xRange = append(xRange, i)
	}

	return xRange
}

func findMissingNumber(x []int) int {
	inputSum := sum(x)

	expectedSum := sum(makeRange(1, len(x)+1))
	return expectedSum - inputSum
}

func main() {
	input := []int{1, 2, 3, 4, 6, 7, 8}

	result := findMissingNumber(input)
	fmt.Println("The missing number is", result)

}
