package rules

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    []string
		wantErr bool
	}{
		{
			name: "empty",
			expr: "",
			want: []string{},
		},
		{
			name: "single",
			expr: "A",
			want: []string{"A"},
		},
		{
			name: "single with spaces",
			expr: "  A  ",
			want: []string{"A"},
		},
		{
			name: "single with tabs",
			expr: "\tA\t",
			want: []string{"A"},
		},
		{
			name: "single with newlines",
			expr: `A AND
B`,
			want: []string{"A", "AND", "B"},
		},
		{
			name: "A AND B",
			expr: "A AND B",
			want: []string{"A", "AND", "B"},
		},
		{
			name: "A AND B AND C",
			expr: "A AND B AND C",
			want: []string{"A", "AND", "B", "AND", "C"},
		},
		{
			name: "NOT A",
			expr: "NOT A",
			want: []string{"NOT", "A"},
		},
		{
			name: "A OR B AND (C EQ D)",
			expr: "A OR B AND (C EQ D)",
			want: []string{"A", "OR", "B", "AND", "(", "C", "EQ", "D", ")"},
		},
		{
			name:    "missing closing parenthesis",
			expr:    "A OR B AND (C EQ D",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tokenize(tt.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   []string
	}{
		{
			name:   "empty",
			tokens: []string{},
			want:   []string{},
		},
		{
			name:   "single",
			tokens: []string{"A"},
			want:   []string{"A"},
		},
		{
			name:   "A AND B",
			tokens: []string{"A", "AND", "B"},
			want:   []string{"A", "B", "AND"},
		},
		{
			name:   "A AND B AND C",
			tokens: []string{"A", "AND", "B", "AND", "C"},
			want:   []string{"A", "B", "AND", "C", "AND"},
		},
		{
			name:   "A OR B AND D EQ C AND C EQ D",
			tokens: []string{"A", "OR", "B", "AND", "D", "EQ", "C", "AND", "C", "EQ", "D"},
			want:   []string{"A", "B", "OR", "D", "C", "EQ", "AND", "C", "D", "EQ", "AND"},
		},
		{
			name:   "NOT A",
			tokens: []string{"NOT", "A"},
			want:   []string{"A", "NOT"},
		},
		{
			name:   "A OR B AND (C EQ D)",
			tokens: []string{"A", "OR", "B", "AND", "(", "C", "EQ", "D", ")"},
			want:   []string{"A", "B", "OR", "C", "D", "EQ", "AND"},
		},
		{
			name:   "A AND B AND (C EQ D)",
			tokens: []string{"A", "AND", "B", "AND", "(", "C", "EQ", "D", ")"},
			want:   []string{"A", "B", "AND", "C", "D", "EQ", "AND"},
		},
		{
			name:   "(C EQ D) AND (E EQ F)",
			tokens: []string{"(", "C", "EQ", "D", ")", "AND", "(", "E", "EQ", "F", ")"},
			want:   []string{"C", "D", "EQ", "E", "F", "EQ", "AND"},
		},
		{
			name:   "A AND (B OR C)",
			tokens: []string{"A", "AND", "(", "B", "OR", "C", ")"},
			want:   []string{"A", "B", "C", "OR", "AND"},
		},
		{
			name:   "A AND B AND (C EQ D) AND (E EQ F)",
			tokens: []string{"A", "AND", "B", "AND", "(", "C", "EQ", "D", ")", "AND", "(", "E", "EQ", "F", ")"},
			want:   []string{"A", "B", "AND", "C", "D", "EQ", "AND", "E", "F", "EQ", "AND"},
		},
		{
			name:   "A AND B AND C LT D",
			tokens: []string{"A", "AND", "B", "AND", "C", "LT", "D"},
			want:   []string{"A", "B", "AND", "C", "D", "LT", "AND"},
		},
		{
			name:   "A AND B AND C EQ D AND E EQ F",
			tokens: []string{"A", "AND", "B", "AND", "C", "EQ", "D", "AND", "E", "EQ", "F"},
			want:   []string{"A", "B", "AND", "C", "D", "EQ", "AND", "E", "F", "EQ", "AND"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.tokens); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
