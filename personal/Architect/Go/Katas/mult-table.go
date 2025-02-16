package main

import "fmt"

func foo() string {
   var str string = ""

   fmt.Println("Start...")
   var number int
   fmt.Print("Please, provide the number: ")
   fmt.Scan(&number)
   
   for i := 1; i <= 10; i++ {
      str += (string(i) + " * " + string(number) + string(i*number) + "\n")
   }
   return str
}

func main() {
   fmt.Print(foo())
}
