package recursion

import "testing"

func CallRecursionTenTimes(t *testing.T) {
	number := 1
	want := 10
	if got := CallRecursionFunc(number); got != want {
		t.Errorf("CallRecursionFunc() = %b, want %b", got, want)
	}
}

func CallRecursionOnce(t *testing.T) {
	number := 15
	want := 15
	if got := CallRecursionFunc(number); got != want {
		t.Errorf("CallRecursionFunc() = %b, want %b", got, want)
	}
}
