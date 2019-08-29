// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsTransferServerInvalidInvocationRoleRule checks the pattern is valid
type AwsTransferServerInvalidInvocationRoleRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsTransferServerInvalidInvocationRoleRule returns new rule with default attributes
func NewAwsTransferServerInvalidInvocationRoleRule() *AwsTransferServerInvalidInvocationRoleRule {
	return &AwsTransferServerInvalidInvocationRoleRule{
		resourceType:  "aws_transfer_server",
		attributeName: "invocation_role",
		pattern:       regexp.MustCompile(`^arn:.*role/.*$`),
	}
}

// Name returns the rule name
func (r *AwsTransferServerInvalidInvocationRoleRule) Name() string {
	return "aws_transfer_server_invalid_invocation_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsTransferServerInvalidInvocationRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsTransferServerInvalidInvocationRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsTransferServerInvalidInvocationRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsTransferServerInvalidInvocationRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`invocation_role does not match valid pattern ^arn:.*role/.*$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
