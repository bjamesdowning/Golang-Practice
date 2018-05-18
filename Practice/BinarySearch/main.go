package main

import "fmt"

///Binary search examples. Recursive search and iterable search.
func main() {
	lst := []int{2, 4, 5, 20, 30, 50, 65, 100, 1000, 1500, 2000, 2456, 2567, 4567, 6779, 34565}
	fmt.Println("RECURSE", recursBinarySearch(lst, 2000, 0, 15))
	fmt.Println("ITER", iterBinarySearch(lst, 1000, 0, 15))
}

func recursBinarySearch(array []int, target, lowIndex, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)

	if array[mid] > target {
		return recursBinarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return recursBinarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target, lowIndex, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] < target {
			startIndex = mid + 1
		} else if array[mid] > target {
			endIndex = mid
		} else {
			return mid
		}
	}
	return -1
}
