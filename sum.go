package main

import "fmt"

func main(){
num := []int{3,3}
fmt.Println(twoSum(num, 6))
}


func twoSum(nums []int, target int) []int {
	idxSlice:= []int{}
	for i, _:= range nums {
		for j, _ := range nums {
			if i != j &&  nums[i] + nums[j] == target{
				idxSlice = append(idxSlice, i)
			}
		}
	}

	return  idxSlice
}