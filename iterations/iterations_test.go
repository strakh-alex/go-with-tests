package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat(5, "a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q, but expected %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Repeat(i, "n")
	}
}

func ExampleRepeat() {
	Repeat(5, "abc")
	// Output: abcabcabcabcabc
}