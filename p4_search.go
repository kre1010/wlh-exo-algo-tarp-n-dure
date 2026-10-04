package main

func SearchV1(ligne []int, v int) int {
	for i := 0; i < len(ligne); i++ {
		if ligne[i] == v {
			return i
		}
	}
	return -1
}
