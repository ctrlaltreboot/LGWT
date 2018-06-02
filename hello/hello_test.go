package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, World... of... tests..."

	if got != want {
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}
