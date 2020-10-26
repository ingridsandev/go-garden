package printfromarray

import "fmt"

// PrintNamesFromArray - print names from array
func PrintNamesFromArray(names ...string) {
	for _, s := range names {
		fmt.Println(s)
	}
}
