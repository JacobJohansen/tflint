// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsBatchComputeEnvironmentInvalidTypeRule checks the pattern is valid
type AwsBatchComputeEnvironmentInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBatchComputeEnvironmentInvalidTypeRule returns new rule with default attributes
func NewAwsBatchComputeEnvironmentInvalidTypeRule() *AwsBatchComputeEnvironmentInvalidTypeRule {
	return &AwsBatchComputeEnvironmentInvalidTypeRule{
		resourceType:  "aws_batch_compute_environment",
		attributeName: "type",
		enum: []string{
			"MANAGED",
			"UNMANAGED",
		},
	}
}

// Name returns the rule name
func (r *AwsBatchComputeEnvironmentInvalidTypeRule) Name() string {
	return "aws_batch_compute_environment_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBatchComputeEnvironmentInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBatchComputeEnvironmentInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBatchComputeEnvironmentInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBatchComputeEnvironmentInvalidTypeRule) Check(runner *tflint.Runner) error {
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
					`type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
