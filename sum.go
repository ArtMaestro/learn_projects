package main

import "fmt"

func maestro(nums []int, target int) (int, int, bool) {
	nMap := make(map[int]int)

	for i, num := range nums {
		comp := target - num
		if j, ok := nMap[comp]; ok {
			return nums[j], num, true
		}
		nMap[num] = i
	}

	return 0, 0, false
}

func main() {
	nums := []int{8, 1, 14, 22}
	target := 0

	num1, num2, ok := maestro(nums, target)

	if ok {
		fmt.Printf("Два числа из массива %v, сумма которых равна %d: %d и %d\n", nums, target, num1, num2)
	} else {
		fmt.Printf("В массиве %v нет двух чисел, сумма которых равна %d\n", nums, target)
	}
}
