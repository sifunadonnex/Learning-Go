package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a simple Go program.")
	fmt.Println("It prints messages to the console.")
	fmt.Println("Goodbye!")
	variableExample()
	stringManipulation()
	sum, diff, prod, quot := arithmeticOperations(10, 5)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quot)
	controlFlowExample()
}

func variableExample() {
	var message string = "This is a variable example."
	var number int = 42
	var isActive bool = true
	const pi float64 = 3.14

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

func arithmeticOperations(a int, b int) (int, int, int, float64) {
	sum := a + b
	diff := a - b
	prod := a * b
	quot := float64(a) / float64(b)
	return sum, diff, prod, quot
}

func controlFlowExample() {
	// Example of control flow with if-else
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
	// Example of control flow with switch
	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	default:
		fmt.Println("Greater than one")
	}
	// Example of control flow with for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	// Example of control flow with range
	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
	// Example of control flow with select
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine"
	}()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("No message received")
	}
	fmt.Println("Control flow example completed.")
}
