package rules

import (
	"strings"
	"unicode"

	"github.com/IAmRadek/rules/internal/utils/stack"
)

// Parse parses a rule expression and returns a Rule.
// The expression is a string that contains a boolean expression.
// The expression can contain the following operators:
//   - AND, OR, XOR, NOT, EQ, NEQ, GT, LT, GTE, LTE
func Parse(name, expr string) (Rule, error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		return nil, ErrEmptyExpression
	}

	output := parse(tokens)

	if !isValid(output) {
		return nil, ErrInvalidExpression
	}

	return &rule{name, output}, nil
}

func isValid(output []string) bool {
	s := make([]bool, 0)

	for _, token := range output {
		if isOperator(token) {
			if token == kNOT {
				if len(s) < 1 {
					return false
				}
				s = s[:len(s)-1]
			} else {
				if len(s) < 2 {
					return false
				}
				s = s[:len(s)-2]
			}
			s = append(s, true)
		} else {
			s = append(s, false)
		}
	}

	return len(s) == 1
}

func MustParse(name, expr string) Rule {
	r, err := Parse(name, expr)
	if err != nil {
		panic(err)
	}
	return r
}

func parse(tokens []string) []string {
	output := make([]string, 0, len(tokens))
	s := stack.Stack[string]{}
	for _, token := range tokens {
		switch token {
		case kAND, kOR, kXOR, kEQ, kNEQ, kGT, kLT, kGTE, kLTE:
			p, ok := s.Peek()
			for ok && precedence[p] >= precedence[token] {
				output = append(output, s.MustPop())
				p, ok = s.Peek()
			}
			s.Push(token)
		case kNOT:
			s.Push(token)
		case "(":
			s.Push(token)
		case ")":
			p, ok := s.Peek()
			for ok && p != "(" {
				output = append(output, s.MustPop())
				p, ok = s.Peek()
			}
			if ok && p == "(" {
				s.MustPop()
			}
			p, ok = s.Peek()
			if ok && p == kNOT {
				output = append(output, s.MustPop())
			}
		default:
			output = append(output, token)
		}
	}
	for !s.IsEmpty() {
		output = append(output, s.MustPop())
	}
	return output
}

var precedence = map[string]int{
	kNOT: 3,
	kAND: 1,
	kOR:  1,
	kXOR: 1,
	kEQ:  2,
	kNEQ: 2,
	kGT:  2,
	kLT:  2,
	kGTE: 2,
	kLTE: 2,
}

func isOperator(char string) bool {
	_, ok := precedence[char]
	return ok
}

func sanitize(expr string) string {
	skipChar := rune(-1)

	mapper := func(r rune) rune {
		if r == '\n' || unicode.IsSpace(r) {
			return ' '
		}
		if r == '\u00A0' {
			return skipChar
		}
		if !unicode.IsDigit(r) && !unicode.IsLetter(r) && r != '(' && r != ')' {
			return skipChar
		}

		return r
	}
	return strings.TrimSpace(strings.Map(mapper, expr))
}

func tokenize(expr string) ([]string, error) {
	expr = sanitize(expr)

	tokens := make([]string, 0, len(expr))
	currentToken := strings.Builder{}
	parenCount := 0

	for _, char := range expr {
		switch char {
		case '(':
			parenCount++
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		case ')':
			parenCount--
			if parenCount < 0 {
				return nil, ErrMismatchedParentheses
			}
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		case ' ':
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
		default:
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	if parenCount != 0 {
		return nil, ErrMismatchedParentheses
	}

	return tokens, nil
}
