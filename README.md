# rules

The `rules` package provides a set of types and functions for creating, parsing, and evaluating boolean
expressions in the form of rules. These rules can be used to define conditions or criteria that determine the
outcome or behavior of a program or system.

## Installation

To install the `rules` package, use `go get`:

```bash
go get github.com/IAmRadek/rules
```

## Usage

Import the rules package in your Go code:

```go
import "github.com/IAmRadek/rules"
```

## Variables

The rules package provides two types of variables: Attribute and Variable.

### `Attribute`

Represents a boolean value that can be used in a rule expression. You can create an attribute using the
NewAttribute function.

```go 
var attr = rules.NewAttribute("attr1")
```

### `Variable`

Represents a dynamic value that can be changed during runtime and used in a rule expression. You can create a
variable using the `NewVariable` function.

```go 
var var1 = rules.NewVariable[float64]("var1")
var var2 = rules.NewVariable[string]("var2")
```

## Rules

The rules package defines a Rule interface that represents a single boolean expression. You can create a rule
using the `Parse` function, which parses a rule expression and returns a `Rule` instance. The expression can
contain operators such as `AND`, `OR`, `XOR`, `NOT`, `EQ`, `NEQ`, `GT`, `LT`, `GTE`, and `LTE`.

``` go
rule, err := rules.Parse("myRule", "var1 AND var2 OR NOT attr1")
if err != nil {
    // handle error
}
```

You can then evaluate a rule using the `Evaluate` method, which takes a `RuleContext` as input and returns a
boolean value and an error indicating whether the rule is true or false.

### RuleContext

The `RuleContext` type holds the values of variables and attributes during the evaluation of rules. You can
create a `RuleContext` using the `NewContext` function, passing in the initial values of variables and
attributes as arguments.

```go 
globalContext := rules.NewContext(var1(10), var2("global"), attr1(true))
userContext := rules.NewContext(var1(12), var2("user")).Merge(globalContext)
```

You can also merge multiple `RuleContext` instances together using the `Merge` method to update the values of
variables and attributes during the evaluation process.

### RuleSet

The `RuleSet` type represents a collection of rules and rule overrides that can be evaluated together as a
group. You can create a RuleSet using the NewRuleSet function, passing in the rules and rule overrides as
arguments.

```go 
ruleSet := rules.NewRuleSet(rule1, rule2, rule3, rule4, rule5, rule6)
```

You can then evaluate the `RuleSet` using the `Evaluate` method, which takes a `RuleContext` as input and
returns a
boolean value and an error indicating whether the rules in the `RuleSet` are true or false.

### RuleOverride

The `RuleOverride` type represents a rule that overrides the behavior of another rule.
You can pass rule to the `AddOverride` method of a `RuleSet` to override the behavior of the rule.

```go
ruleSet.AddOverride(rule1)
```

### Contributing

If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a
pull request on GitHub at https://github.com/IAmRadek/rules/issues.

