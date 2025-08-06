package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints messages to the console.")
	fmt.Println("Goodbye!")
	variableExample()
	stringManipulation()
}

func variableExample() {
	var message string = "This is a variable example."
	var number int = 42
	var isActive bool = true
	var pi float64 = 3.14

	fmt.Println(message)
	fmt.Println("Number:", number)
	fmt.Println("Is Active:", isActive)
	fmt.Println("Value of Pi:", pi)
	fmt.Println("Variable example completed.")
}

func stringManipulation() {
	var str string = "Hello, Go!" + " Welcome to string manipulation."
	fmt.Println(str)
	fmt.Println("Length of string:", len(str))
	fmt.Println("Substring:", str[7:9])
	fmt.Println("String in uppercase:", "HELLO, GO!")
	fmt.Println("String manipulation completed.")
}
