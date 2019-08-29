// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudformationStackInvalidTemplateURLRule checks the pattern is valid
type AwsCloudformationStackInvalidTemplateURLRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackInvalidTemplateURLRule returns new rule with default attributes
func NewAwsCloudformationStackInvalidTemplateURLRule() *AwsCloudformationStackInvalidTemplateURLRule {
	return &AwsCloudformationStackInvalidTemplateURLRule{
		resourceType:  "aws_cloudformation_stack",
		attributeName: "template_url",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackInvalidTemplateURLRule) Name() string {
	return "aws_cloudformation_stack_invalid_template_url"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackInvalidTemplateURLRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackInvalidTemplateURLRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackInvalidTemplateURLRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackInvalidTemplateURLRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"template_url must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"template_url must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
