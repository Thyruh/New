package main

import "fmt"

func is_bool(n int) bool {
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
	for i := 2; i <= 10000; i++ {
	 	if is_bool(i) {
			fmt.Printf("The number %v is Prime\n", i)
		}
	}
}




