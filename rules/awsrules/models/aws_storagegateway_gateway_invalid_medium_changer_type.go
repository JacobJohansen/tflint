// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayGatewayInvalidMediumChangerTypeRule checks the pattern is valid
type AwsStoragegatewayGatewayInvalidMediumChangerTypeRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsStoragegatewayGatewayInvalidMediumChangerTypeRule returns new rule with default attributes
func NewAwsStoragegatewayGatewayInvalidMediumChangerTypeRule() *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule {
	return &AwsStoragegatewayGatewayInvalidMediumChangerTypeRule{
		resourceType:  "aws_storagegateway_gateway",
		attributeName: "medium_changer_type",
		max:           50,
		min:           2,
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Name() string {
	return "aws_storagegateway_gateway_invalid_medium_changer_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayGatewayInvalidMediumChangerTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"medium_changer_type must be 50 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"medium_changer_type must be 2 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
