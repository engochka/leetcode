package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 3, 8}
	fmt.Printf("Число отличающихся %d\n", removeElement(nums, 3))
}

func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	//fmt.Printf("Массив после сортировки:%d\n", nums)
	return k
}
