package rules

import (
	"fmt"
)

// Variable represents a value that can be used in a rule.
type Variable interface {
	RuleElement

	getValue() any

	equalTo(Variable) Attribute
	notEqualTo(Variable) Attribute
	greaterThan(Variable) Attribute
	lessThan(Variable) Attribute
	greaterThanOrEqualTo(Variable) Attribute
	lessThanOrEqualTo(Variable) Attribute
}

type variable[T any] struct {
	name  string
	value T

	eq func(v1, v2 T) bool
	gt func(v1, v2 T) bool
}

type variableFunc[T any] func(value T) variable[T]

func (v variableFunc[T]) getType() string {
	return "variable"
}

func (v variableFunc[T]) getName() string {
	var zero T
	return v(zero).getName()
}

type ordered interface {
	~float32 | ~float64 | ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~string
}

func NewVariable[T ordered](name string) variableFunc[T] {
	return func(value T) variable[T] {
		return variable[T]{
			name:  name,
			value: value,
			eq:    func(v1, v2 T) bool { return v1 == v2 },
			gt:    func(v1, v2 T) bool { return v1 > v2 },
		}
	}
}

func (v variable[T]) String() string {
	return fmt.Sprintf("%v(%v)", v.name, v.value)
}

func (v variable[T]) getType() string {
	return "variable"
}

func (v variable[T]) getName() string {
	return v.name
}

func (v variable[T]) getValue() any {
	return v.value
}

func (v variable[T]) equalTo(v2 Variable) Attribute {
	a1 := v.getValue().(T)
	a2 := v2.getValue().(T)
	if v.eq(a1, a2) {
		return attribute{name: "(" + v.name + " == " + v2.getName() + ")", value: true}
	}
	return attribute{name: "(" + v.name + " != " + v2.getName() + ")", value: false}
}

func (v variable[T]) notEqualTo(v2 Variable) Attribute {
	return v.equalTo(v2).not()
}

func (v variable[T]) greaterThan(v2 Variable) Attribute {
	a1 := v.getValue().(T)
	a2 := v2.getValue().(T)
	if v.gt(a1, a2) {
		name := "(" + v.name + " > " + v2.getName() + ")"
		return attribute{name: name, value: true}
	}
	name := "(" + v.name + " <= " + v2.getName() + ")"
	return attribute{name: name, value: false}
}

func (v variable[T]) lessThan(v2 Variable) Attribute {
	return v.greaterThan(v2).not()
}

func (v variable[T]) greaterThanOrEqualTo(v2 Variable) Attribute {
	return v.greaterThan(v2).or(v.equalTo(v2))
}

func (v variable[T]) lessThanOrEqualTo(v2 Variable) Attribute {
	return v.lessThan(v2).or(v.equalTo(v2))
}
