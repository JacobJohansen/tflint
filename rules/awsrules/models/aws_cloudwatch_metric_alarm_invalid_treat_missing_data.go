// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudwatchMetricAlarmInvalidTreatMissingDataRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidTreatMissingDataRule() *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule {
	return &AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "treat_missing_data",
		max:           255,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_treat_missing_data"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidTreatMissingDataRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"treat_missing_data must be 255 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"treat_missing_data must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
