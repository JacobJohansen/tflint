// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSfnStateMachineInvalidRoleArnRule checks the pattern is valid
type AwsSfnStateMachineInvalidRoleArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSfnStateMachineInvalidRoleArnRule returns new rule with default attributes
func NewAwsSfnStateMachineInvalidRoleArnRule() *AwsSfnStateMachineInvalidRoleArnRule {
	return &AwsSfnStateMachineInvalidRoleArnRule{
		resourceType:  "aws_sfn_state_machine",
		attributeName: "role_arn",
		max:           256,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSfnStateMachineInvalidRoleArnRule) Name() string {
	return "aws_sfn_state_machine_invalid_role_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSfnStateMachineInvalidRoleArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSfnStateMachineInvalidRoleArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSfnStateMachineInvalidRoleArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSfnStateMachineInvalidRoleArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"role_arn must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"role_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
