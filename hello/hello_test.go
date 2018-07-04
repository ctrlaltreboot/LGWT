package main

import "testing"

/*
	  *testing.T is our hook into the testing framework...
		type t exposes functions for stuff we might want as reaction
		or feedback on our tests..
*/
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got: '%s', Want: '%s'", got, want)
		}
	}

	t.Run("Say hello to handsome people", func(t *testing.T) {
		got := Hello("handsome")
		want := "Hello there, handsome"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to the world when there's an empty argument", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
