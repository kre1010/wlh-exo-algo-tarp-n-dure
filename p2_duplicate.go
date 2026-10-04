package main

func DuplicateV1(pile []int) int {
	for i := 0; i < len(pile); i++ {
		for j := i + 1; j < len(pile); j++ {
			if pile[i] == pile[j] {
				return pile[i]
			}
		}
	}
	return -1
}

func DuplicateV2(pile []int) int {
	vus := make(map[int]bool)

	for i := 0; i < len(pile); i++ {
		if vus[pile[i]] {
			return pile[i]
		}
		vus[pile[i]] = true
	}
	return -1
}
