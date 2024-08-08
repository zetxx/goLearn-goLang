package main

func Sum(list []int) int {
	var sum int = 0
	for _, num := range list {
		sum = sum + num
	}
	return sum
}

func SumAll(sumLists ...[]int) (sum []int) {
	for _, list := range sumLists {
		sum = append(sum, Sum(list))
	}
	return
}

func SumAllTails(sumLists ...[]int) (sum []int) {
	for _, list := range sumLists {
		sum = append(sum, Sum(list[1:]))
	}
	return
}
