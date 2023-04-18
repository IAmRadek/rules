package rules

import (
	"fmt"
	"strings"
)

type ruleContext []RuleElement

func (r *ruleContext) String() string {
	s := strings.Builder{}
	for i, elem := range *r {
		if i == len(*r)-1 {
			s.WriteString(fmt.Sprint(elem))
			continue
		}
		s.WriteString(fmt.Sprint(elem) + ", ")
	}

	return s.String()
}

func NewContext(elems ...RuleElement) RuleContext {
	var ctx ruleContext
	ctx = append(ctx, elems...)

	return &ctx
}

func (r *ruleContext) listElements() []RuleElement {
	return *r
}

// MergeWith combines two contexts into a new context. If an element with the same
// name exists in both contexts, the element from the first context is used.
func (r *ruleContext) MergeWith(ctx RuleContext) RuleContext {
	newCtx := ruleContext{}
	newCtx = append(newCtx, *r...)
	for _, elem := range ctx.listElements() {
		if _, ok := r.findElement(elem.getName()); !ok {
			newCtx = append(newCtx, elem)
		}
	}

	return &newCtx
}

func (r *ruleContext) findElement(name string) (RuleElement, bool) {
	for _, elem := range *r {
		if elem.getName() == name {
			return elem, true
		}
	}
	return nil, false
}
