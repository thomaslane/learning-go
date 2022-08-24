package main

import "fmt"

func main() {
	uArr := []int32{5, 7, 9, 22, 13, 15, 1, 0, 2}
	fmt.Println("unsorted...")
	fmt.Println(uArr)
	quickSort(uArr, 0, len(uArr)-1)
	fmt.Println("sorted...")
	fmt.Println(uArr)
}

// func quickSort(arr []int32, start, end int) []int32 {
// 	fmt.Printf("Calling quickSort(%v, %v, %v)\n", arr, start, end)
// 	if start < end {
// 		fmt.Println(start, "is less than", end)
// 		var pivot int
// 		arr, pivot = partition(arr, start, end)
// 		arr = quickSort(arr, start, pivot-1)
// 		arr = quickSort(arr, pivot+1, end)
// 	}
// 	fmt.Printf("Quicksort returning %v\n", arr)
// 	return arr
// }

// func partition(arr []int32, start, end int) ([]int32, int) {
// 	pivot := arr[end]
// 	splitIndex := start
// 	fmt.Printf("calling partition(%v, %v, %v)\n", arr, start, end)
// 	fmt.Println("pivot:", pivot, "i:", splitIndex)
// 	for j := start; j < end; j++ {
// 		fmt.Printf("j: %v\tarr[j]: %v\ti: %v\tarr[i]: %v\n", j, arr[j], splitIndex, arr[splitIndex])
// 		if arr[j] < pivot {
// 			fmt.Printf("%v is less than %v\n", arr[j], pivot)
// 			fmt.Println("Swapping", arr[splitIndex], arr[j])
// 			arr[splitIndex], arr[j] = arr[j], arr[splitIndex]
// 			splitIndex++
// 		}
// 	}
// 	fmt.Println("End of loop, swapping", arr[splitIndex], arr[end])
// 	arr[splitIndex], arr[end] = arr[end], arr[splitIndex]
// 	return arr, splitIndex
// }

func quickSort(arr []int32, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			arr[i], arr[splitIndex] = arr[splitIndex], arr[i]
			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	quickSort(arr, start, splitIndex-1)
	quickSort(arr, splitIndex+1, end)
}
