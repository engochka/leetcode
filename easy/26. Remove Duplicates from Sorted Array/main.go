package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	k := 1
	j := 1
	i := 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i+1] = nums[j]
			i++
			k++
			j++
		}
	}
	fmt.Printf("%#v\n", nums)
	return k
}
