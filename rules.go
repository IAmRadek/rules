/*
Package rules provides a set of tools for parsing, evaluating, and working with boolean expressions in the form of rules.
It includes support for variables, attributes, and rule sets.
*/
package rules

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidRule is an error indicating that a rule is invalid.
	ErrInvalidRule = fmt.Errorf("invalid rule")
	// ErrMissingDataInContext is an error indicating that there is a lack of data in the rule context.
	ErrMissingDataInContext = fmt.Errorf("missing data in context")
	// ErrMismatchedParentheses is an error indicating that there are mismatched parentheses in a rule expression.
	ErrMismatchedParentheses = errors.New("mismatched parentheses")
	// ErrEmptyExpression is an error indicating that the rule expression is empty.
	ErrEmptyExpression = errors.New("empty expression")
	// ErrInvalidExpression is an error indicating that the rule expression is invalid.
	ErrInvalidExpression = errors.New("invalid expression")
)

// RuleElement is an interface that represents a rule element, which can be an attribute, a variable, or any other element of a rule.
type RuleElement interface {
	getType() string
	getName() string
}

// Rule is an interface that represents a rule.
// It includes methods for getting the name of the rule and evaluating the rule with a given rule context to determine its boolean result.
type Rule interface {
	Name() string
	Evaluate(ctx RuleContext) (bool, error)
}

type RuleSet interface {
	AddRule(rule Rule)
	AddOverride(override RuleOverride)
	Evaluate(ctx RuleContext) (bool, error)
}

type RuleOverride interface {
	Name() string
}

// RuleContext is an interface that represents a rule context, which is a collection of rule elements.
// It includes methods for merging two rule contexts.
type RuleContext interface {
	Merge(ctx RuleContext) RuleContext

	findElement(name string) (RuleElement, bool)
	listElements() []RuleElement
}
