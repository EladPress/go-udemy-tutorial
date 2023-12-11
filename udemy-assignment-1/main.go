package main

import "fmt"

func main() {
	nums := []int{}

	for i := 0; i <= 10; i++ {
		nums = append(nums, i)
	}

	for _, number := range nums {
		if number%2 == 0 {
			fmt.Println(number, "even")
		} else {
			fmt.Println(number, "odd")
		}
	}

	// fmt.Println(nums)
}
