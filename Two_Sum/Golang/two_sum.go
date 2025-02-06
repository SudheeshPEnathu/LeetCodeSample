package main

import "fmt"

func main() {
	fmt.Println("Script to find the two sum problem")

	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSumS1(nums, target)
	fmt.Println("Resulr 1 - ", result)

	nums = []int{3, 2, 4}
	target = 6
	result = twoSumS2(nums, target)
	fmt.Println("Resulr 1 - ", result)

	nums = []int{3, 3}
	target = 6
	result = twoSumS3(nums, target)
	fmt.Println("Resulr 1 - ", result)

	nums = []int{3, 3}
	target = 7
	result = twoSumS3(nums, target)
	fmt.Println("Resulr 1 - ", result)
}

func twoSumS1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumS2(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		numsMap[num] = i
	}

	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok && i != j {
			return []int{i, j}
		}
	}

	return nil
}

func twoSumS3(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numsMap[target-num]; ok && i != j {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return nil
}
