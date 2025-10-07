package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkSub(b *testing.B) {
	b.Run("Wahid", func(b *testing.B) {
		for b.Loop() {
			HelloWorld("Wahid")
		}
	})
	b.Run("Rizka", func(b *testing.B) {
		for b.Loop() {
			HelloWorld("Rizka")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	// old way
	for i := 0; i < b.N; i++ {
		HelloWorld("Wahid")
	}

	// modernize
	// for b.Loop() {
	// 	HelloWorld("Wahid")
	// }
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Wahid",
			request:  "Wahid",
			expected: "Hello Wahid",
		},
		{
			name:     "Rizka",
			request:  "Rizka",
			expected: "Hello Rizka",
		},
		{
			name:     "Fathur",
			request:  "Fathur",
			expected: "Hello Fathur",
		},
		{
			name:     "Drian",
			request:  "Drian",
			expected: "Hello Drian",
		},
		{
			name:     "Fajar",
			request:  "Fajar",
			expected: "Hello Fajar",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Wahid", func(t *testing.T) {
		result := HelloWorld("Wahid")
		require.Equal(t, "Hello Wahid", result, "Result must be 'Hello Wahid'")
	})

	t.Run("Rizka", func(t *testing.T) {
		result := HelloWorld("Rizka")
		require.Equal(t, "Hello Rizka", result, "Result must be 'Hello Rizka'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

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
