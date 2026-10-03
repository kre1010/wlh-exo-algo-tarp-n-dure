package main 

func TowerHeightV1(n int) int {
	var aaa int = 0 
	for i := 0; i >= n ; i++ {
		aaa += i
	}
	return aaa
}