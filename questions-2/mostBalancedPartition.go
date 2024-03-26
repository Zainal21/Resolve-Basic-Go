package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'mostBalancedPartition' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY parent
 *  2. INTEGER_ARRAY files_size
 */

func mostBalancedPartition(parent []int32, files_size []int32) int32 {
	//Get sum at each node.
	size := make(map[int32]int32, len(parent))
	var total int32 = 0

	//add own size
	for i := range files_size {
		size[int32(i)] = files_size[i]
		//if we cut i from parent
	}

	for i := len(parent) - 1; i >= 0; i-- {
		p := parent[i]

		s := size[p]
		size[p] = size[int32(i)] + s
		total = total + files_size[i]
		//new size at parent now
	}

	minDiff := total //there is no cut yet
	for i := range files_size {
		nodeSize := size[int32(i)]
		other := total - nodeSize
		diff := int32(math.Abs(float64(other - nodeSize)))
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	parent := []int32{0, 1, 1, 2, 2}
	files_size := []int32{4, 6, 2, 6, 1}

	result := mostBalancedPartition(parent, files_size)
	fmt.Println(result)
}
