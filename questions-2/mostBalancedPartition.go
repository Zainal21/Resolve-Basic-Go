package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
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
	// Write your code here
	// Write your code here
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
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	parentCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var parent []int32

	for i := 0; i < int(parentCount); i++ {
		parentItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		parentItem := int32(parentItemTemp)
		parent = append(parent, parentItem)
	}

	files_sizeCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var files_size []int32

	for i := 0; i < int(files_sizeCount); i++ {
		files_sizeItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		files_sizeItem := int32(files_sizeItemTemp)
		files_size = append(files_size, files_sizeItem)
	}

	result := mostBalancedPartition(parent, files_size)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
