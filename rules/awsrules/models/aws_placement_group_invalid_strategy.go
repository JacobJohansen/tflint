// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsPlacementGroupInvalidStrategyRule checks the pattern is valid
type AwsPlacementGroupInvalidStrategyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsPlacementGroupInvalidStrategyRule returns new rule with default attributes
func NewAwsPlacementGroupInvalidStrategyRule() *AwsPlacementGroupInvalidStrategyRule {
	return &AwsPlacementGroupInvalidStrategyRule{
		resourceType:  "aws_placement_group",
		attributeName: "strategy",
		enum: []string{
			"cluster",
			"spread",
			"partition",
		},
	}
}

// Name returns the rule name
func (r *AwsPlacementGroupInvalidStrategyRule) Name() string {
	return "aws_placement_group_invalid_strategy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsPlacementGroupInvalidStrategyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsPlacementGroupInvalidStrategyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsPlacementGroupInvalidStrategyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsPlacementGroupInvalidStrategyRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

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
					`strategy is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
