package main 

import "fmt"

const (
   monday = iota + 1
   tuesday
   wednesday
   thursday
   friday
   saturday
   sunday
)

func main() { 
   fmt.Printf("monday:  %v\n", monday)
   fmt.Printf("tuesday: %v\n", tuesday)
   fmt.Printf("wednesday: %v\n", wednesday)
   fmt.Printf("thursday: %v\n", thursday)
   fmt.Printf("friday: %v\n", friday)
   fmt.Printf("saturday: %v\n", saturday)
   fmt.Printf("sunday: %v\n", sunday)
}
