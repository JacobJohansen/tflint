// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferUserInvalidUserNameRule checks the pattern is valid
type AwsTransferUserInvalidUserNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsTransferUserInvalidUserNameRule returns new rule with default attributes
func NewAwsTransferUserInvalidUserNameRule() *AwsTransferUserInvalidUserNameRule {
	return &AwsTransferUserInvalidUserNameRule{
		resourceType:  "aws_transfer_user",
		attributeName: "user_name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9_][a-zA-Z0-9_-]{2,31}$`),
	}
}

// Name returns the rule name
func (r *AwsTransferUserInvalidUserNameRule) Name() string {
	return "aws_transfer_user_invalid_user_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferUserInvalidUserNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferUserInvalidUserNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferUserInvalidUserNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferUserInvalidUserNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`user_name does not match valid pattern ^[a-zA-Z0-9_][a-zA-Z0-9_-]{2,31}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
