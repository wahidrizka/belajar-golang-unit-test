package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}

	result := HelloWorld("Wahid")
	require.Equal(t, "Hello Wahid", result, "Result must be 'Hello Wahid'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Wahid")
	require.Equal(t, "Hello Wahid", result, "Result must be 'Hello Wahid'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Wahid")
	assert.Equal(t, "Hello Wahid", result, "Result must be 'Hello Wahid'")
	fmt.Println("TestHelloWorldAssert Done")
}

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
		t.Fatal()
	}

	fmt.Println("TestHelloWorldFathur Done")
}
