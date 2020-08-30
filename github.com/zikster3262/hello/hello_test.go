package main

import "testing"

func assertCorrectMessageString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David!"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David", "")
		want := "Hello, David!"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David", "Czech")
		want := "Ahoj, David!"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David", "English")
		want := "Hello, David!"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David", "Vietnamese")
		want := "Xian chao, David!"

		assertCorrectMessageString(t, got, want)
	})
}
