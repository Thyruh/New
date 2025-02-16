package main

import "fmt"


func is_prime(n int) bool {
   if n <= 2 {
      return false
   }
   for i := 2; i*i <= n; i++ {
      if n % i == 0 {
         return false
      }
   }
   return true
}


func main() {
   var num int
   for num < 100000 {
      if is_prime(num) {
         fmt.Printf("The number %v is prime.\n", num)
      } 
      num++
   }
   
}
