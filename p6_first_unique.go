package main

func FirstUniqueV1(ligne []int) int {
	compteur := make(map[int]int)
	for _, kapla := range ligne {
		compteur[kapla]++
	}
	for _, kapla := range ligne {
		if compteur[kapla] == 1 {
			return kapla
		}
	}
	return -1
}
