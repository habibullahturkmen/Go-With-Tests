package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	assertRepeat := func(t testing.TB, repeated, expected string) {
		if repeated != expected {
			t.Errorf("expected %q, but got %q", expected, repeated)
		}
	}

	t.Run("repeat five times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertRepeat(t, repeated, expected)
	})

	t.Run("repeat fifteen times", func(t *testing.T) {
		repeated := Repeat("b", 15)
		expected := "bbbbbbbbbbbbbbb"
		assertRepeat(t, repeated, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeat2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x", 20)
	}
}

func BenchmarkRepeat3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("abc", 30)
	}
}

func ExampleRepeat() {
	repeated := Repeat("abc", 3)
	fmt.Println(repeated)
	// Output: abcabcabc
}