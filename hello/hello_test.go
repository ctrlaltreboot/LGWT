package main

import "testing"

/*
	  *testing.T is our hook into the testing framework

		type t exposes functions for stuff we might want as reaction
		or feedback on our tests..
*/
func TestHello(t *testing.T) {

	// declare function and assign it to a variable
	// question: what's the benefit of assigning a function to a variable versus defining an outside function?
	// clue: scoping: memory allocation bit...
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper() // tell the test suite that am a helper function
		if got != want {
			t.Errorf("Got: '%s', Want: '%s'", got, want)
		}
	}

	t.Run("Say hello to handsome people", func(t *testing.T) {
		got := Hello("handsome", "")
		want := "Hello, handsome"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to the world when there's an empty argument", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Beautiful"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "french")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
