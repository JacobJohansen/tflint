// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferSSHKeyInvalidBodyRule checks the pattern is valid
type AwsTransferSSHKeyInvalidBodyRule struct {
	resourceType  string
	attributeName string
	max           int
	pattern       *regexp.Regexp
}

// NewAwsTransferSSHKeyInvalidBodyRule returns new rule with default attributes
func NewAwsTransferSSHKeyInvalidBodyRule() *AwsTransferSSHKeyInvalidBodyRule {
	return &AwsTransferSSHKeyInvalidBodyRule{
		resourceType:  "aws_transfer_ssh_key",
		attributeName: "body",
		max:           2048,
		pattern:       regexp.MustCompile(`^ssh-rsa\s+[A-Za-z0-9+/]+[=]{0,3}(\s+.+)?\s*$`),
	}
}

// Name returns the rule name
func (r *AwsTransferSSHKeyInvalidBodyRule) Name() string {
	return "aws_transfer_ssh_key_invalid_body"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferSSHKeyInvalidBodyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferSSHKeyInvalidBodyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferSSHKeyInvalidBodyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferSSHKeyInvalidBodyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"body must be 2048 characters or less",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`body does not match valid pattern ^ssh-rsa\s+[A-Za-z0-9+/]+[=]{0,3}(\s+.+)?\s*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
