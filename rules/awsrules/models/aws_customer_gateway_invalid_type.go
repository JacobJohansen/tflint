// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCustomerGatewayInvalidTypeRule checks the pattern is valid
type AwsCustomerGatewayInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCustomerGatewayInvalidTypeRule returns new rule with default attributes
func NewAwsCustomerGatewayInvalidTypeRule() *AwsCustomerGatewayInvalidTypeRule {
	return &AwsCustomerGatewayInvalidTypeRule{
		resourceType:  "aws_customer_gateway",
		attributeName: "type",
		enum: []string{
			"ipsec.1",
		},
	}
}

// Name returns the rule name
func (r *AwsCustomerGatewayInvalidTypeRule) Name() string {
	return "aws_customer_gateway_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCustomerGatewayInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCustomerGatewayInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCustomerGatewayInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCustomerGatewayInvalidTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
