package main

import (
	"bufio"
	"os"
	greet "petunia/greeting"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	println("Enter your name:")
	name, _ := reader.ReadString('\n')

	response := greet.SayHello(name)

	println("Response:", response)
}
