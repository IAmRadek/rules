package rules

import (
	"fmt"

	"github.com/IAmRadek/rules/internal/utils/stack"
)

type rule struct {
	name string
	r    []string
}

func (r *rule) Name() string {
	return r.name
}

func (r *rule) String() string {
	st := stack.Stack[string]{}
	for i := 0; i < len(r.r); i++ {
		token := r.r[i]
		if token == kNOT {
			op1 := st.MustPop()
			st.Push(token + " " + op1)
			continue
		}
		if precedence[token] > 0 {
			op2 := st.MustPop()
			op1 := st.MustPop()

			// Determine whether to add parentheses
			needsParens := false
			if i < len(r.r)-1 && precedence[r.r[i+1]] > precedence[token] {
				needsParens = false
			}

			if needsParens {
				st.Push("(" + op1 + " " + token + " " + op2 + ")")
			} else {
				st.Push(op1 + " " + token + " " + op2)
			}
		} else {
			st.Push(token)
		}
	}
	return st.MustPop()
}

func (r *rule) Evaluate(ctx RuleContext) (bool, error) {
	i := 0

	st := &stack.Stack[RuleElement]{}
	for {
		if i >= len(r.r) {
			break
		}
		op := r.r[i]
		switch op {
		case kNOT:
			s1, ok := st.Pop()
			if !ok {
				return false, fmt.Errorf("%w: missing operand for NOT operator", ErrInvalidRule)
			}
			s1a, ok := s1.(Attribute)
			if !ok {
				return false, fmt.Errorf("%w: operand for NOT operator must be an attribute", ErrInvalidRule)
			}

			st.Push(s1a.not())
			i++
			continue
		case kAND:
			s1, s2, err := popTwoAttributes(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for AND operator", ErrInvalidRule)
			}

			st.Push(s2.and(s1))
		case kOR:
			s1, s2, err := popTwoAttributes(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for OR operator", ErrInvalidRule)
			}
			st.Push(s2.or(s1))
		case kXOR:
			s1, s2, err := popTwoAttributes(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for XOR operator", ErrInvalidRule)
			}
			st.Push(s2.xor(s1))
		case kEQ:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for EQ operator", ErrInvalidRule)
			}
			st.Push(s2.equalTo(s1))
		case kNEQ:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for NEQ operator", ErrInvalidRule)
			}
			st.Push(s2.notEqualTo(s1))
		case kGT:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for GT operator", ErrInvalidRule)
			}
			st.Push(s2.greaterThan(s1))
		case kLT:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for LT operator", ErrInvalidRule)
			}
			st.Push(s2.lessThan(s1))
		case kGTE:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for GTE operator", ErrInvalidRule)
			}
			st.Push(s2.greaterThanOrEqualTo(s1))
		case kLTE:
			s1, s2, err := popTwoVariables(st)
			if err != nil {
				return false, fmt.Errorf("%w: missing operand for LTE operator", ErrInvalidRule)
			}
			st.Push(s2.lessThanOrEqualTo(s1))
		default:
			el, ok := ctx.findElement(op)
			if !ok {
				return false, fmt.Errorf("%w: %s", ErrMissingDataInContext, op)
			}
			switch v := el.(type) {
			case Attribute:
				st.Push(v)
			case Variable:
				st.Push(v)
			}
		}
		i++
	}

	if out, ok := st.Pop(); ok {
		return out.(Attribute).getValue(), nil
	}

	return false, fmt.Errorf("%w: no output attribute", ErrInvalidRule)
}

func popTwoAttributes(st *stack.Stack[RuleElement]) (Attribute, Attribute, error) {
	s1, ok := st.Pop()
	if !ok {
		return nil, nil, fmt.Errorf("%w: no output attribute", ErrInvalidRule)
	}
	s2, ok := st.Pop()
	if !ok {
		return nil, nil, fmt.Errorf("%w: no output attribute", ErrInvalidRule)
	}

	s1a, ok := s1.(Attribute)
	if !ok {
		return nil, nil, fmt.Errorf("%w, expected attribute, got %T", ErrInvalidRule, s1)
	}

	s2a, ok := s2.(Attribute)
	if !ok {
		return nil, nil, fmt.Errorf("%w, expected attribute, got %T", ErrInvalidRule, s2)
	}

	return s1a, s2a, nil
}

func popTwoVariables(st *stack.Stack[RuleElement]) (Variable, Variable, error) {
	s1, ok := st.Pop()
	if !ok {
		return nil, nil, fmt.Errorf("%w: no output attribute", ErrInvalidRule)
	}
	s2, ok := st.Pop()
	if !ok {
		return nil, nil, fmt.Errorf("%w: no output attribute", ErrInvalidRule)
	}

	s1a, ok := s1.(Variable)
	if !ok {
		return nil, nil, fmt.Errorf("%w: expected variable, got %T", ErrInvalidRule, s1)
	}

	s2a, ok := s2.(Variable)
	if !ok {
		return nil, nil, fmt.Errorf("%w: expected variable, got %T", ErrInvalidRule, s2)
	}

	return s1a, s2a, nil
}
