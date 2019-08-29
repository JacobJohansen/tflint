// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoUserPoolClientInvalidUserPoolIDRule checks the pattern is valid
type AwsCognitoUserPoolClientInvalidUserPoolIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolClientInvalidUserPoolIDRule returns new rule with default attributes
func NewAwsCognitoUserPoolClientInvalidUserPoolIDRule() *AwsCognitoUserPoolClientInvalidUserPoolIDRule {
	return &AwsCognitoUserPoolClientInvalidUserPoolIDRule{
		resourceType:  "aws_cognito_user_pool_client",
		attributeName: "user_pool_id",
		max:           55,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w-]+_[0-9a-zA-Z]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolClientInvalidUserPoolIDRule) Name() string {
	return "aws_cognito_user_pool_client_invalid_user_pool_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolClientInvalidUserPoolIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolClientInvalidUserPoolIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolClientInvalidUserPoolIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolClientInvalidUserPoolIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"user_pool_id must be 55 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user_pool_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`user_pool_id does not match valid pattern ^[\w-]+_[0-9a-zA-Z]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
