// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsDmsEndpointInvalidSslModeRule checks the pattern is valid
type AwsDmsEndpointInvalidSslModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDmsEndpointInvalidSslModeRule returns new rule with default attributes
func NewAwsDmsEndpointInvalidSslModeRule() *AwsDmsEndpointInvalidSslModeRule {
	return &AwsDmsEndpointInvalidSslModeRule{
		resourceType:  "aws_dms_endpoint",
		attributeName: "ssl_mode",
		enum: []string{
			"none",
			"require",
			"verify-ca",
			"verify-full",
		},
	}
}

// Name returns the rule name
func (r *AwsDmsEndpointInvalidSslModeRule) Name() string {
	return "aws_dms_endpoint_invalid_ssl_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDmsEndpointInvalidSslModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDmsEndpointInvalidSslModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDmsEndpointInvalidSslModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDmsEndpointInvalidSslModeRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

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
					`ssl_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
