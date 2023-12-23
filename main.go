package main

import "fmt"

func main() {
	list := []int{1, 3, 4, 5, 7, 9}
	target := 5
	index := binarySearch(list, target)

	if index == -1 {
		fmt.Println("Index not found")
		return
	}

	output := fmt.Sprintf("The index is %d", index)
	fmt.Println(output)
}

// Get the position
func binarySearch(list []int, target int) int {
	startIndex := 0
	endIndex := len(list) - 1

	for startIndex <= endIndex {
		middleIndex := (startIndex + endIndex) / 2
		middleNumber := list[middleIndex]

		if middleNumber == target {
			return middleIndex
		}

		if middleNumber < target {
			startIndex = middleIndex + 1
		} else {
			endIndex = middleIndex - 1
		}

	}

	return -1
}
