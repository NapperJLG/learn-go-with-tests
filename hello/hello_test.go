package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectTesting := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jake")
		want := "Hello Jake"

		assertCorrectTesting(t, got, want)
	})

	t.Run("saying Hello Stranger if empty string supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello Stranger"

		assertCorrectTesting(t, got, want)
	})

}
