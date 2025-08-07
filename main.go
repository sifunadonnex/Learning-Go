package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// main function serves as the entry point of the program
// It demonstrates various Go features including variables, strings, arithmetic operations,
// control flow, arrays, slices, maps, runes, and pointers.

var wg sync.WaitGroup
var dbData = []string{"data1", "data2", "data3"}

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
	arraysandSlicesExample()
	mapsExample()
	runesExample()
	pointerExample()

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("All DB calls completed in:", time.Since(t0))
	fmt.Println("Program completed successfully.")

	// Example of a goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine iteration:", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	channelExample()
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

func arraysandSlicesExample() {
	// Example of arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println("Length of array:", len(arr))
	fmt.Println("First element:", arr[0])

	// Example of slices
	var slice []int = []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("First element:", slice[0])
	fmt.Println("Last element:", slice[len(slice)-1])
	fmt.Println("Array and slice example completed.")
}

func mapsExample() {
	// Example of maps
	var m map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map:", m)
	fmt.Println("Value for key 'two':", m["two"])
	delete(m, "two")
	fmt.Println("Map after deletion:", m)
	fmt.Println("Maps example completed.")
}

func runesExample() {
	var mystring string = "Hello, 世界"
	fmt.Println("String:", mystring)
	fmt.Println("Length of string in bytes:", len(mystring))
	fmt.Println("Length of string in runes:", len([]rune(mystring)))
	for i, r := range mystring {
		fmt.Printf("Character %d: %c\n", i, r)
	}
	fmt.Println("Runes example completed.")
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("Happy birthday %s! You are now %d.\n", p.Name, p.Age)
}

func pointerExample() {
	var p *int32 = new(int32)
	*p = 42
	fmt.Println("Pointer value:", *p)
	var person Person = Person{Name: "Alice", Age: 30}
	person.Greet()
	person.HaveBirthday()
	fmt.Println("Pointer example completed.")
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2.0
	time.Sleep(time.Duration(delay * float32(time.Second)))
	fmt.Printf("the result of db call %d is %s\n", i, dbData[i%len(dbData)])
	wg.Done()
}

func channelExample() {
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
	fmt.Println("Channel example completed.")
}
