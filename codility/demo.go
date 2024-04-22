//Write a function:
//
//    func Solution(A []int) int
//
//that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.
//
//For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

//Given A = [1, 2, 3], the function should return 4.

//Given A = [−1, −3], the function should return 1.

//Write an efficient algorithm for the following assumptions:

//        N is an integer within the range [1..100,000];
//        each element of array A is an integer within the range [−1,000,000..1,000,000].

//

package main

import (
	"fmt"
)

func Solution(A []int) int {
	// Initialize a map to track which positive integers are present
	positiveSet := make(map[int]bool)

	// Populate the map with positive integers from the array
	for _, value := range A {
		if value > 0 {
			positiveSet[value] = true
		}
	}

	// Start checking from 1 to find the smallest missing positive integer
	for i := 1; i <= len(A)+1; i++ {
		if !positiveSet[i] {
			return i
		}
	}

	// This line should technically never be reached because the for loop includes len(A)+1
	return len(A) + 1
}

func main() {
	// Test cases
	fmt.Println(Solution([]int{1, 3, 6, 4, 1, 2})) // Expected output: 5
	fmt.Println(Solution([]int{1, 2, 3}))          // Expected output: 4
	fmt.Println(Solution([]int{-1, -3}))           // Expected output: 1
}
