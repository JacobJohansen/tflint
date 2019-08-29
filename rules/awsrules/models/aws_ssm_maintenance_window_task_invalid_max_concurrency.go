// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule checks the pattern is valid
type AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule returns new rule with default attributes
func NewAwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule() *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule {
	return &AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule{
		resourceType:  "aws_ssm_maintenance_window_task",
		attributeName: "max_concurrency",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[1-9][0-9]%|[1-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule) Name() string {
	return "aws_ssm_maintenance_window_task_invalid_max_concurrency"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmMaintenanceWindowTaskInvalidMaxConcurrencyRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"max_concurrency must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"max_concurrency must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`max_concurrency does not match valid pattern ^([1-9][0-9]*|[1-9][0-9]%|[1-9]%|100%)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
