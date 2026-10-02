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




