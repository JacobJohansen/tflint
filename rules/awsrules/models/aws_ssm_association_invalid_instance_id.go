// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmAssociationInvalidInstanceIDRule checks the pattern is valid
type AwsSsmAssociationInvalidInstanceIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidInstanceIDRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidInstanceIDRule() *AwsSsmAssociationInvalidInstanceIDRule {
	return &AwsSsmAssociationInvalidInstanceIDRule{
		resourceType:  "aws_ssm_association",
		attributeName: "instance_id",
		pattern:       regexp.MustCompile(`^(^i-(\w{8}|\w{17})$)|(^mi-\w{17}$)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidInstanceIDRule) Name() string {
	return "aws_ssm_association_invalid_instance_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidInstanceIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidInstanceIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidInstanceIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidInstanceIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`instance_id does not match valid pattern ^(^i-(\w{8}|\w{17})$)|(^mi-\w{17}$)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
