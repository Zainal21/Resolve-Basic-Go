package main

import (
	"fmt"
)

/*
 * Complete the 'findSum' function below.
 *
 * The function is expected to return a LONG_INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY numbers
 *  2. 2D_INTEGER_ARRAY queries
 */

func findSum(numbers []int32, queries [][]int32) []int64 {
	resultArray := []int64{}
	for i := 0; i < len(queries); i++ {
		start := queries[i][0] - 1
		end := queries[i][1] - 1
		zeroValue := queries[i][2]

		result := int64(0)
		for k := start; k <= end; k++ {
			targetNum := numbers[k]
			if targetNum == 0 {
				result += int64(zeroValue)
			} else {
				result += int64(targetNum)
			}
		}
		resultArray = append(resultArray, result)
	}

	return resultArray
}

func main() {
	numbers := []int32{5, 10, 10, 20}
	queries := [][]int32{{1, 2, 5}}
	result := findSum(numbers, queries)
	fmt.Println(result)
}
