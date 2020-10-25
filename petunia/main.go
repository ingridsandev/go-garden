package main

import (
	"bufio"
	"os"
	greet "petunia/greeting"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	println("Enter your name:")
	name, _ := reader.ReadString('\n')

	response := greet.SayHello(strings.TrimSuffix(name, "\n"))

	println("Response:", response)
}
