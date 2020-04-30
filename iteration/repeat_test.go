package iteration

import (
	"fmt"
	"testing"
)

func TestRepeater(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Got %s, Expected: %s ", repeated, expected)
	}
}

func ExampleRepeat() {
	r := Repeat("x")
	fmt.Println(r)
	// Output: xxxxx
}
