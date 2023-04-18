package rules

type ruleSet struct {
	rules     map[string]Rule
	overrides []RuleOverride
}

func (r *ruleSet) AddRule(rule Rule) {
	r.rules[rule.Name()] = rule
}

func (r *ruleSet) AddOverride(override RuleOverride) {
	r.overrides = append(r.overrides, override)
}

func (r *ruleSet) Evaluate(ctx RuleContext) (bool, error) {
	for _, rule := range r.rules {
		if r.isOverridden(rule) {
			continue
		}

		result, err := rule.Evaluate(ctx)
		if err != nil {
			return false, err
		}
		if !result {
			return false, nil
		}
	}
	return true, nil
}

func (r *ruleSet) isOverridden(r2 Rule) bool {
	for _, override := range r.overrides {
		if override.Name() == r2.Name() {
			return true
		}
	}
	return false
}

func NewRuleSet(rules ...Rule) RuleSet {
	rs := &ruleSet{
		rules: make(map[string]Rule),
	}
	for _, rule := range rules {
		rs.AddRule(rule)
	}
	return rs
}
