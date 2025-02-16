package main

import (
	"fmt"
	"unicode"
)


func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func asciiSum(word string) int {
	sum := 0
	for _, ch := range word {
		if unicode.IsLetter(ch) {
			sum += int(ch)
		}
	}
	return sum
}


func main() {
	var word string;
	fmt.Print("Enter a word: ")
	fmt.Scan(&word)

	sum := asciiSum(word)
	fmt.Printf("The ASCII sum of the word '%s' is %d.\n", word, sum)

	if isPrime(sum) {
		fmt.Printf("The word '%s' is PRIME!\n", word)
	} else {
		fmt.Printf("The word '%s' is NOT PRIME.\n", word)
	}
}




































