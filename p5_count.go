package main

func CountV1(pile []int, plafond int) []int {
	count := make([]int, plafond+1)
	for _, value := range pile {
		count[value]++
	}
	return count
}
