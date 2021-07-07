/*
Given an n x n matrix where each of the rows and columns are sorted in ascending order, return the kth smallest element in the matrix.
Note that it is the kth smallest element in the sorted order, not the kth distinct element.



Example 1:
Input:
	matrix = [[1,5,9],[10,11,13],[12,13,15]], k = 8
Output:
	13
Explanation:
	The elements in the matrix are [1,5,9,10,11,12,13,13,15], and the 8th smallest number is 13

Example 2:
Input:
	matrix = [[-5]], k = 1
Output:
	-5

Constraints:
    n == matrix.length
    n == matrix[i].length
    1 <= n <= 300
    -109 <= matrix[i][j] <= 109
    All the rows and columns of matrix are guaranteed to be sorted in non-decreasing order.
    1 <= k <= n2

*/

package goleetcode

import "sort"

//Notes: The solution provided below works but is very hacky in nature. Instead of taking advantage of the sorted matrix, I flattened it and sorted it
// This is highly inefficient and there should be an iterative way to improve on this. I will take a look into it afterwards

func kthSmallest(matrix [][]int, k int) int {
	var flat []int
	for i := 0; i < len(matrix); i++ {
		flat = append(flat, matrix[i]...)
	}

	sort.Ints(flat)

	return flat[k-1]
}
