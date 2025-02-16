package main

import "fmt"

func main() {
	var limit int = 100
	for i := 1; i <= limit; i++ {
		if i > 1 {
			fmt.Printf(", ")
		}
		fmt.Print(i)		
	}
	fmt.Println(".")
}

// type masturbation
