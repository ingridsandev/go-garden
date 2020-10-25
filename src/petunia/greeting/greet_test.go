package greet

import "testing"

func TestHello(t *testing.T) {
	name := "Ingrid"
	want := "Hello, " + name
	if got := SayHello(name); got != want {
		t.Errorf("SayHello() = %q, want %q", got, want)
	}
}
