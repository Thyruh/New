package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."
	var userLog string = fmt.Sprintf("%v %v, Age: %v, Rate: %.10f, Is subscribed: %v, Message: %v",
	fname, lname, age, messageRate, isSubscribed, message)

	// Don't touch above this line

	// ?

	// Don't touch below this line

	fmt.Println(userLog)
}

// Name: FNAME LNAME, Age: AGE, Rate: MESSAGERATE, Is Subscribed: ISSUBSCRIBED, Message: MESSAGE

/* Where FNAME LNAME AGE MESSAGERATE ISSUBSCRIBED and MESSAGE correspond to the variables above.

   MESSAGERATE should be rounded to the tenths place.
*/
