// compute the Fibonacci series for a given n

// questions
// should the function return the number in the sequence that a specific number is
// or should it return all of the numbers in the sequence up to that point?

package main

import "fmt"

type found map[int]uint64

// 0 1 2 3 4

func calcFibNum(n int, m found) uint64 {
	if _, ok := m[n]; ok {
		return m[n]
	}
	if n == 0 {
		m[n] = 0
		return 0
	} else if n == 1 {
		m[n] = 1
		return 1
	} else {
		num := calcFibNum(n-2, m) + calcFibNum(n-1, m)
		m[n] = num
		return num
	}
}

func main() {
	num := 100
	fibNums := make(found)
	calcFibNum(num, fibNums)
	fmt.Println(fibNums)
}
