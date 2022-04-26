package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if expected != repeated {
			t.Errorf("expected %q but got %q\n", expected, repeated)
		}
	}
	t.Run("repeating characters a fixed time", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("repeating a charater with variable time", func(t *testing.T) {
		repeated := Repeat("x", 3)
		expected := "xxx"

		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 3)
	}
}

func ExampleRepeat() {
	repeated := Repeat("x", 5)
	fmt.Println(repeated)
	// Output = "xxxxx"
}
