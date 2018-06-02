package main

import "testing"

// *testing.T is our hook into the testing framework...
func TestHello(t *testing.T) {
	got := Hello("handsome")
	want := "Hello there, handsome"

	if got != want {
		/*
			type t exposes functions for stuff we might want as reaction
			or feedback on our tests..
		*/
		t.Errorf("Got: '%s', Want: '%s'", got, want)
	}
}
