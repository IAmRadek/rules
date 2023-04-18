package rules

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	tests := []struct {
		rule string
	}{
		{
			rule: "A AND B",
		},
		{
			rule: "A AND B AND C",
		},
		{
			rule: "NOT A",
		},
		{
			rule: "A OR B AND B EQ C AND C EQ D",
		},
		{
			rule: "A AND B AND C EQ D",
		},
		{
			rule: "C EQ D AND E EQ F",
		},
		{
			rule: "A AND B OR C",
		},
		{
			rule: "A AND B AND C EQ D AND E EQ F",
		},
		{
			rule: "A AND B AND C LT D",
		},
		{
			rule: "A AND B AND C EQ D AND E EQ F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.rule, func(t *testing.T) {
			r1, err := Parse("rule", tt.rule)
			if err != nil {
				t.Fatal(err)
			}

			stringed1 := fmt.Sprintf("%s", r1)

			r2, err := Parse("rule", stringed1)
			if err != nil {
				t.Fatal(err)
			}

			stringed2 := fmt.Sprintf("%s", r2)
			if stringed1 != stringed2 {
				t.Errorf("String() = %v, want %v", stringed1, stringed2)
			}

			if stringed1 != tt.rule {
				t.Errorf("String() = %v, want %v", stringed1, tt.rule)
			}
		})
	}
}
