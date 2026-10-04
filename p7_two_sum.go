package main

func TwoSumV1(pile []int, cible int) (int, int, bool) {
	for i := 0; i < len(pile); i++ {
		for j := i + 1; j < len(pile); j++ {
			if pile[i]+pile[j] == cible {
				return pile[i], pile[j], true
			}
		}
	}
	return 0, 0, false
}
