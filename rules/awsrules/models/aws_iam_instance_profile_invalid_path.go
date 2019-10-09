// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsIAMInstanceProfileInvalidPathRule checks the pattern is valid
type AwsIAMInstanceProfileInvalidPathRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMInstanceProfileInvalidPathRule returns new rule with default attributes
func NewAwsIAMInstanceProfileInvalidPathRule() *AwsIAMInstanceProfileInvalidPathRule {
	return &AwsIAMInstanceProfileInvalidPathRule{
		resourceType:  "aws_iam_instance_profile",
		attributeName: "path",
		max:           512,
		min:           1,
		pattern:       regexp.MustCompile(`^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`),
	}
}

// Name returns the rule name
func (r *AwsIAMInstanceProfileInvalidPathRule) Name() string {
	return "aws_iam_instance_profile_invalid_path"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMInstanceProfileInvalidPathRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMInstanceProfileInvalidPathRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMInstanceProfileInvalidPathRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMInstanceProfileInvalidPathRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"path must be 512 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"path must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`path does not match valid pattern ^(\x{002F})|(\x{002F}[\x{0021}-\x{007F}]+\x{002F})$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
