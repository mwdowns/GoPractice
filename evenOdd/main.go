package main

import "fmt"

type numSlice = []int

func newNumSlice(num int) numSlice {
	nums := numSlice{}
	for i := 0; i <= num; i++ {
		nums = append(nums, i)
	}
	return nums
}

func main() {
	nums := newNumSlice(10)
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
