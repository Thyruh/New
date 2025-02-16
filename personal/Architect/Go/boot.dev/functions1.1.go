package main

import "fmt"

func getMonthlyPrice(tier string) int {
	switch tier {
	case "basic":
		return 100 * 100
	case "premium":
		return 150 * 100
	case "enterprise":
		return 500 * 100
	}
	return 0
}


func main() {
	result1 := getMonthlyPrice("basic") 
	result2 := getMonthlyPrice("enterprise")
	fmt.Println(result1, result2)
}
