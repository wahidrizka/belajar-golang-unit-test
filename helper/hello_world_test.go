package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Wahid")

	if result != "Hello Wahid" {
		// error
		panic("Result is not 'Hello Wahid'")
	}
}

func TestHelloWorldFathur(t *testing.T) {
	result := HelloWorld("Fathur")

	if result != "Hello Fathur" {
		// error
		panic("Result is not 'Hello Fathur'")
	}
}
