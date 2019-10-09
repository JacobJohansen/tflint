// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsWafregionalRegexPatternSetInvalidNameRule checks the pattern is valid
type AwsWafregionalRegexPatternSetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsWafregionalRegexPatternSetInvalidNameRule returns new rule with default attributes
func NewAwsWafregionalRegexPatternSetInvalidNameRule() *AwsWafregionalRegexPatternSetInvalidNameRule {
	return &AwsWafregionalRegexPatternSetInvalidNameRule{
		resourceType:  "aws_wafregional_regex_pattern_set",
		attributeName: "name",
		max:           128,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsWafregionalRegexPatternSetInvalidNameRule) Name() string {
	return "aws_wafregional_regex_pattern_set_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafregionalRegexPatternSetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafregionalRegexPatternSetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafregionalRegexPatternSetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafregionalRegexPatternSetInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
