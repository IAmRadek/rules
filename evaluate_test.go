package rules

import (
	"testing"
)

func TestEvaluate(t *testing.T) {
	var A = NewAttribute("A")
	var B = NewAttribute("B")
	var C = NewVariable[string]("C")
	var D = NewVariable[string]("D")

	tests := []struct {
		rule    string
		ctx     RuleContext
		wantErr bool
	}{
		{
			rule: "A",
			ctx:  NewContext(A(true)),
		},
		{
			rule: "NOT A",
			ctx:  NewContext(A(false)),
		},
		{
			rule: "A AND B",
			ctx:  NewContext(A(true), B(true)),
		},
		{
			rule: "A OR B",
			ctx:  NewContext(A(true), B(false)),
		},
		{
			rule: "A XOR B",
			ctx:  NewContext(A(true), B(false)),
		},
		{
			rule: "A AND NOT B",
			ctx:  NewContext(A(true), B(false)),
		},
		{
			rule: "A OR NOT B",
			ctx:  NewContext(A(true), B(false)),
		},
		{
			rule: "A AND B AND C LT D",
			ctx:  NewContext(A(true), B(true), C("C"), D("D")),
		},
		{
			rule: "A AND B AND C GT D",
			ctx:  NewContext(A(true), B(true), C("D"), D("C")),
		},
		{
			rule: "A AND B AND C EQ D AND C EQ D",
			ctx:  NewContext(A(true), B(true), C("D"), D("D")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.rule, func(t *testing.T) {
			got, err := Parse("rule", tt.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			result, err := got.Evaluate(tt.ctx)
			if err != nil {
				t.Fatal(err)
			}

			if !result {
				t.Errorf("Parse() got = %v, want %v", got, tt.rule)
			}
		})
	}
}
