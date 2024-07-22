package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

// function dengan parameter
func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

// function dengan return value
func getHello(firstname string) string{
	return "Hello " + firstname
}

// function dengan multiple return value
func getFullName(test string) (string, string, int){
	return "Ibnu", test, 15
}

// Kode program named return values
func getCompleteName() (firstName, middleName, lastName string){
	firstName = "Ibnu"
	middleName = "Jarir"
	lastName = "Ath Thobari"

	return firstName, middleName, lastName
}

// kode program variadic function
func sumAll(numbers ...int) int{
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

//kode program function value
func getGoodBye(name string) string{
	return "Good bye " + name 
}

// kode program function as parameter dan function type declarations
type Filter func(string) string
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// Anonymous function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// kode program factorial for loop
func factorialLoop(value int) int{
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//kode program factorial recursive
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}


func main1() {
	sayHello()
	sayHelloTo("Agaton", "Dev")
	result := getHello("Agaton")
	fmt.Println(result)
	firstName, lastName, _ := getFullName("Umar")
	fmt.Println(firstName, lastName)
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
	
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)
	//slice parameter
	numbers := []int{10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)

	goodBye := getGoodBye
	fmt.Println(goodBye("Malas"))

	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("eko", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))

	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)

}