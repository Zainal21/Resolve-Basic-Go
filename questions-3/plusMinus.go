package main

import "fmt"

func plusMinus(arr []int32) {
	pos, neg, zero := 0, 0, 0

	for _, num := range arr {
		if num > 0 {
			pos++
		} else if num == 0 {
			zero++
		} else {
			neg++
		}
	}

	size := float32(len(arr))
	fmt.Println(float32(pos) / size)
	fmt.Println(float32(neg) / size)
	fmt.Println(float32(zero) / size)
}

func main() {
	numbers := []int32{-1, 2, 0, -3, 4, 0}
	plusMinus(numbers)
}
