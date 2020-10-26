package recursion

import "strconv"

// CallRecursionFunc - Call itself til number is > than 10
func RecursionFunc(number int) int {
	println("started n-" + strconv.Itoa(number))

	if number > 10 {
		println(strconv.Itoa(number) + " is > than 10 - exiting")
		return number
	}

	println(strconv.Itoa(number) + " is <= than 10")
	number++
	return RecursionFunc(number)
}
