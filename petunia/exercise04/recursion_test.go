package recursion

import "testing"

func TestCallRecursionTenTimes(t *testing.T) {
	number := 1
	want := 11
	if got := CallRecursionFunc(number); got != want {
		t.Errorf("CallRecursionFunc() = %d, want %d", got, want)
	}
}

func TestCallRecursionOnce(t *testing.T) {
	number := 15
	want := 15
	if got := CallRecursionFunc(number); got != want {
		t.Errorf("CallRecursionFunc() = %d, want %d", got, want)
	}
}
