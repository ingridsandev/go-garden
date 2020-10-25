package checknumberinbetween

import "testing"

func TestCheckNumberIsInBetween_true(t *testing.T) {
	number := 1
	initial := 1
	end := 10
	want := true
	if got := CheckNumberIsInBetween(number, initial, end); got != want {
		t.Errorf("CheckNumberIsInBetween() = %t, want %t", got, want)
	}
}

func TestCheckNumberIsInBetween_false(t *testing.T) {
	number := 11
	initial := 1
	end := 10
	want := false
	if got := CheckNumberIsInBetween(number, initial, end); got != want {
		t.Errorf("CheckNumberIsInBetween() = %t, want %t", got, want)
	}
}
