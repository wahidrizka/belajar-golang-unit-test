package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldWahid(t *testing.T) {
	result := HelloWorld("Wahid")

	if result != "Hello Wahid" {
		// error
		t.Error("Result must be 'Hello Wahid'")
	}

	fmt.Println("TestHelloWorldWahid Done")
}

func TestHelloWorldFathur(t *testing.T) {
	result := HelloWorld("Fathur")

	if result != "Hello Fathur" {
		// error
		t.Fatal("Result must be 'Hello Fathur'")
	}

	fmt.Println("TestHelloWorldFathur Done")
}
