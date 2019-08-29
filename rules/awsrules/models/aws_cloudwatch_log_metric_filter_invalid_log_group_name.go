// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule checks the pattern is valid
type AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchLogMetricFilterInvalidLogGroupNameRule returns new rule with default attributes
func NewAwsCloudwatchLogMetricFilterInvalidLogGroupNameRule() *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule {
	return &AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule{
		resourceType:  "aws_cloudwatch_log_metric_filter",
		attributeName: "log_group_name",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^[\.\-_/#A-Za-z0-9]+$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule) Name() string {
	return "aws_cloudwatch_log_metric_filter_invalid_log_group_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchLogMetricFilterInvalidLogGroupNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"log_group_name must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"log_group_name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`log_group_name does not match valid pattern ^[\.\-_/#A-Za-z0-9]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
