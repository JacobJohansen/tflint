// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmMaintenanceWindowTaskInvalidTaskArnRule checks the pattern is valid
type AwsSsmMaintenanceWindowTaskInvalidTaskArnRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsSsmMaintenanceWindowTaskInvalidTaskArnRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTaskInvalidTaskArnRule() *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule {
	return &AwsSsmMaintenanceWindowTaskInvalidTaskArnRule{
		resourceType:  "aws_ssm_maintenance_window_task",
		attributeName: "task_arn",
		max:           1600,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule) Name() string {
	return "aws_ssm_maintenance_window_task_invalid_task_arn"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTaskInvalidTaskArnRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"task_arn must be 1600 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"task_arn must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
