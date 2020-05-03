package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {

	assertCorrectRepeats := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("Got: %s, Want: %s", got, want)
		}
	}

	t.Run("Repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectRepeats(t, got, want)
	})

	t.Run("Repeat 10 times", func(t *testing.T) {
		got := Repeat("x", 10)
		want := "xxxxxxxxxx"
		assertCorrectRepeats(t, got, want)
	})
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	r := Repeat("x", 5)
	fmt.Println(r)
	// Output: xxxxx
}
