package main

import (
	greet "GoGarden/petunia/exercise01"
	checknumberinbetween "GoGarden/petunia/exercise02"
	printfromarray "GoGarden/petunia/exercise03"
	recursion "GoGarden/petunia/exercise04"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Exercise 01
	//CallGreetingMethod()

	// Exercise 02
	//CallCheckNumberIsInBetweenMethod()

	// Exercise 03
	//CallPrintNamesFromArray()

	// Exercise 04
	//CallRecursionFunc()
}

func CallRecursionFunc()  {
	response := recursion.RecursionFunc(1)
	println("Response:", response)
}

// CallPrintNamesFromArray - Call CallPrintNamesFromArray method
func CallPrintNamesFromArray() {
	a := []string{"Alice", "Bob", "Charley"}
	printfromarray.PrintNamesFromArray(a...)
}

// CallGreetingMethod - Call greeting method
func CallGreetingMethod() {
	reader := bufio.NewReader(os.Stdin)

	println("Enter your name:")
	name, _ := reader.ReadString('\n')

	response := greet.SayHello(strings.TrimSuffix(name, "\n"))

	println("Response:", response)
}

// CallCheckNumberIsInBetweenMethod - Call CallCheckNumberIsInBetweenMethod method
func CallCheckNumberIsInBetweenMethod() {
	reader := bufio.NewReader(os.Stdin)

	println("Using ; as delimitator, enter a number, an initial value and an end value to compate if number is in between - eg: 1;10;0")
	value, _ := reader.ReadString('\n')

	// remove new line and split using ; as its delimitator
	values := strings.Split(strings.TrimSuffix(value, "\n"), ";")

	// Convert first item to int
	number, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatal(err)
	}

	// Convert second item to int
	initial, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatal(err)
	}

	// Convert third item to int
	end, err := strconv.Atoi(values[2])
	if err != nil {
		log.Fatal(err)
	}

	response := checknumberinbetween.CheckNumberIsInBetween(number, initial, end)

	println("Response:", response)
}
