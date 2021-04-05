package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	cases := []struct {
		character string
		times     int
		expected  string
	}{
		{"a", 1, "a"},
		{"a", 2, "aa"},
		{"a", 3, "aaa"},
		{"a", 4, "aaaa"},
		{"a", 5, "aaaaa"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets repeated %d times to %q", test.character, test.times, test.expected), func(t *testing.T) {
			got := Repeat(test.character, test.times)

			if got != test.expected {
				t.Errorf("expected %q but got %q", test.expected, got)
			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
