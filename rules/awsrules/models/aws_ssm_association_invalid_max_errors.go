// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsSsmAssociationInvalidMaxErrorsRule checks the pattern is valid
type AwsSsmAssociationInvalidMaxErrorsRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsSsmAssociationInvalidMaxErrorsRule returns new rule with default attributes
func NewAwsSsmAssociationInvalidMaxErrorsRule() *AwsSsmAssociationInvalidMaxErrorsRule {
	return &AwsSsmAssociationInvalidMaxErrorsRule{
		resourceType:  "aws_ssm_association",
		attributeName: "max_errors",
		max:           7,
		min:           1,
		pattern:       regexp.MustCompile(`^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`),
	}
}

// Name returns the rule name
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Name() string {
	return "aws_ssm_association_invalid_max_errors"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmAssociationInvalidMaxErrorsRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"max_errors must be 7 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"max_errors must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`max_errors does not match valid pattern ^([1-9][0-9]*|[0]|[1-9][0-9]%|[0-9]%|100%)$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
