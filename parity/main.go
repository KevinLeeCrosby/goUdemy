package main

import "fmt"

func main() {
   //integers := []int{}
   integers := make([]int, 0)

   for i := 0; i < 11; i++ {
      integers = append(integers, i)
   }

   for _, integer := range integers {
      checkParity(integer)
   }
}

func checkParity(i int) {
   if i & 1 == 1 {
      fmt.Println(i, "is odd")
   } else {
      fmt.Println(i, "is even")
   }
}
