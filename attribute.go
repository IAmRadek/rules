package rules

import (
	"fmt"
)

// Attribute represents a boolean value that can be used in a rule.
type Attribute interface {
	RuleElement

	getValue() bool

	and(Attribute) Attribute
	or(Attribute) Attribute
	xor(Attribute) Attribute
	not() Attribute
}

type attribute struct {
	name  string
	value bool
}

type AttributeDescriptor func(value bool) Attribute

func (v AttributeDescriptor) getType() string {
	return "attribute"
}

func (v AttributeDescriptor) getName() string {
	return v(false).getName()
}

func NewAttribute(name string) AttributeDescriptor {
	return func(value bool) Attribute {
		return attribute{
			name:  name,
			value: value,
		}
	}
}

func (p attribute) String() string {
	return fmt.Sprintf("%s(%t)", p.name, p.value)
}

func (p attribute) getType() string {
	return "attribute"
}

func (p attribute) getName() string {
	return p.name
}

func (p attribute) getValue() bool {
	return p.value
}

func (p attribute) and(p2 Attribute) Attribute {
	return attribute{value: p.value && p2.getValue()}
}

func (p attribute) or(p2 Attribute) Attribute {
	return attribute{value: p.value || p2.getValue()}
}

func (p attribute) xor(p2 Attribute) Attribute {
	return attribute{value: p.value != p2.getValue()}
}

func (p attribute) not() Attribute {
	return attribute{value: !p.value}
}
