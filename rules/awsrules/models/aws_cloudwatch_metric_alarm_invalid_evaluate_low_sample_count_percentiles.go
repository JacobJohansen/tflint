// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule() *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule {
	return &AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "evaluate_low_sample_count_percentiles",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_evaluate_low_sample_count_percentiles"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidEvaluateLowSampleCountPercentilesRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"evaluate_low_sample_count_percentiles must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"evaluate_low_sample_count_percentiles must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
