package greet

// SayHello - Greet someone by name as an argument
func SayHello(name string) string {
	var especialName = "Ingrid"

	if name == especialName {
		return "Hello, Coxis"

	}
	return "Hello, " + name
}
