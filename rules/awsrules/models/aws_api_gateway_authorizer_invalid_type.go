// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsAPIGatewayAuthorizerInvalidTypeRule checks the pattern is valid
type AwsAPIGatewayAuthorizerInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAPIGatewayAuthorizerInvalidTypeRule returns new rule with default attributes
func NewAwsAPIGatewayAuthorizerInvalidTypeRule() *AwsAPIGatewayAuthorizerInvalidTypeRule {
	return &AwsAPIGatewayAuthorizerInvalidTypeRule{
		resourceType:  "aws_api_gateway_authorizer",
		attributeName: "type",
		enum: []string{
			"TOKEN",
			"REQUEST",
			"COGNITO_USER_POOLS",
		},
	}
}

// Name returns the rule name
func (r *AwsAPIGatewayAuthorizerInvalidTypeRule) Name() string {
	return "aws_api_gateway_authorizer_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAPIGatewayAuthorizerInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAPIGatewayAuthorizerInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAPIGatewayAuthorizerInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAPIGatewayAuthorizerInvalidTypeRule) Check(runner *tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
