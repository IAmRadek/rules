package rules

import (
	"fmt"
	"testing"
)

func FuzzParse(f *testing.F) {
	f.Add("A AND B AND C LT D")
	f.Add("A AND B AND C GT D")
	f.Add("A AND B AND C EQ D")
	f.Add("A AND B AND C NE D")
	f.Add("A AND B AND C LE D")
	f.Add("A AND B AND C GE D")
	f.Add("A AND B AND C EQ D AND E EQ F")
	f.Add("A AND B AND (C EQ D) AND (E EQ F)")
	f.Add("A AND (B OR C)")
	f.Add("(C EQ D) AND (E EQ F)")
	f.Add("A AND B AND (C EQ D)")
	f.Add("A OR B AND (C EQ D)")
	f.Add("NOT A")
	f.Add("A OR B AND D EQ C AND C EQ D")
	f.Add("A AND B AND C")
	f.Add("A AND B")
	f.Add("A")
	f.Add("A AND B AND (C EQ D) AND (E EQ F)")

	f.Fuzz(func(t *testing.T, b string) {
		r1, err := Parse("rule", b)
		if err != nil {
			return
		}
		s1 := fmt.Sprint(r1)

		r2, err := Parse("rule", s1)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		s2 := fmt.Sprint(r2)
		if s1 != s2 {
			t.Fatalf("%q, expected %q, got %q", b, s1, s2)
		}
	})
}
