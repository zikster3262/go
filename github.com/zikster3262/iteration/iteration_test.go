package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := Repeat("a", 15)
	want := "aaaaaaaaaaaaaaa"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 15)
	}
}

func ExampleRepeat() {
	Repeat("a", 15)
	// Output: aaaaaaaaaaaaaaa
}
