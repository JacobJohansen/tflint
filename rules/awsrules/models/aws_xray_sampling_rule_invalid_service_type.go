// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsXraySamplingRuleInvalidServiceTypeRule checks the pattern is valid
type AwsXraySamplingRuleInvalidServiceTypeRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsXraySamplingRuleInvalidServiceTypeRule returns new rule with default attributes
func NewAwsXraySamplingRuleInvalidServiceTypeRule() *AwsXraySamplingRuleInvalidServiceTypeRule {
	return &AwsXraySamplingRuleInvalidServiceTypeRule{
		resourceType:  "aws_xray_sampling_rule",
		attributeName: "service_type",
		max:           64,
	}
}

// Name returns the rule name
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Name() string {
	return "aws_xray_sampling_rule_invalid_service_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsXraySamplingRuleInvalidServiceTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"service_type must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
