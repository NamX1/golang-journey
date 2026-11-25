package main

import "fmt"

// Declared and initialized: message starts with this value.
var message string = "Hello, World!"

// Declared without an explicit value; Go initializes it to an empty string.
var message2 string

var number int = 42      // An integer initialized to 42.
var pi float64 = 3.14159 // A floating-point number initialized to 3.14159.
var isActive bool = true // A boolean initialized to true.

func main() {
	fmt.Println(message) // Prints: Hello, World!

	message2 = "Today is a great day!" // Assigns a value to the already-declared variable.
	fmt.Println(message2)              // Prints: Today is a great day!

	fmt.Println("Number:", number)      // Prints: Number: 42
	fmt.Println("Pi:", pi)              // Prints: Pi: 3.14159
	fmt.Println("Is Active:", isActive) // Prints: Is Active: true
}
