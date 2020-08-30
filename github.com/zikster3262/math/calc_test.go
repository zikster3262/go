package main

import "testing"

func assertCorrectMessage(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessageString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertCorrectMessageFloat(t *testing.T, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %b want %b", got, want)
	}
}

func TestCalc(t *testing.T) {
	t.Run("run with values", func(t *testing.T) {
		got := Calc(5, 13)
		want := 18

		assertCorrectMessage(t, got, want)
	})
}

func TestDivide(t *testing.T) {
	t.Run("run Divide", func(t *testing.T) {
		got := Divide(5, 12)
		want := 0.4166666666666667

		assertCorrectMessageFloat(t, got, want)
	})
}

func TestHello(t *testing.T) {
	t.Run("run Hello", func(t *testing.T) {
		got := Hello("David")
		want := "Hello David"

		assertCorrectMessageString(t, got, want)
	})

	t.Run("run Hello", func(t *testing.T) {
		got := Hello("Peter")
		want := "Hello Peter"

		assertCorrectMessageString(t, got, want)
	})
}
