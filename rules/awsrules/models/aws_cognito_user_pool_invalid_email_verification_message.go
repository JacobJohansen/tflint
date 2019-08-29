// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCognitoUserPoolInvalidEmailVerificationMessageRule checks the pattern is valid
type AwsCognitoUserPoolInvalidEmailVerificationMessageRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule returns new rule with default attributes
func NewAwsCognitoUserPoolInvalidEmailVerificationMessageRule() *AwsCognitoUserPoolInvalidEmailVerificationMessageRule {
	return &AwsCognitoUserPoolInvalidEmailVerificationMessageRule{
		resourceType:  "aws_cognito_user_pool",
		attributeName: "email_verification_message",
		max:           20000,
		min:           6,
		pattern:       regexp.MustCompile(`^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Name() string {
	return "aws_cognito_user_pool_invalid_email_verification_message"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoUserPoolInvalidEmailVerificationMessageRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"email_verification_message must be 20000 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"email_verification_message must be 6 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`email_verification_message does not match valid pattern ^[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*\{####\}[\p{L}\p{M}\p{S}\p{N}\p{P}\s*]*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
