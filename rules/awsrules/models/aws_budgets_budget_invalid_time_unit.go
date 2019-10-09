// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsBudgetsBudgetInvalidTimeUnitRule checks the pattern is valid
type AwsBudgetsBudgetInvalidTimeUnitRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBudgetsBudgetInvalidTimeUnitRule returns new rule with default attributes
func NewAwsBudgetsBudgetInvalidTimeUnitRule() *AwsBudgetsBudgetInvalidTimeUnitRule {
	return &AwsBudgetsBudgetInvalidTimeUnitRule{
		resourceType:  "aws_budgets_budget",
		attributeName: "time_unit",
		enum: []string{
			"DAILY",
			"MONTHLY",
			"QUARTERLY",
			"ANNUALLY",
		},
	}
}

// Name returns the rule name
func (r *AwsBudgetsBudgetInvalidTimeUnitRule) Name() string {
	return "aws_budgets_budget_invalid_time_unit"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBudgetsBudgetInvalidTimeUnitRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBudgetsBudgetInvalidTimeUnitRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBudgetsBudgetInvalidTimeUnitRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBudgetsBudgetInvalidTimeUnitRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`time_unit is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
