package main

import "fmt"

// myMessage prints the same message whenever it is called.
func myMessage() {
	fmt.Println("This is a message from the function.")
}

// greet uses its string parameter to personalize the greeting.
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	myMessage()

	// Each call runs myMessage again.
	myMessage()
	myMessage()

	greet("Kevin") // Pass "Kevin" as the name argument.
}
