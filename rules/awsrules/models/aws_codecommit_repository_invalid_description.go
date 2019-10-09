// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsCodecommitRepositoryInvalidDescriptionRule checks the pattern is valid
type AwsCodecommitRepositoryInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsCodecommitRepositoryInvalidDescriptionRule returns new rule with default attributes
func NewAwsCodecommitRepositoryInvalidDescriptionRule() *AwsCodecommitRepositoryInvalidDescriptionRule {
	return &AwsCodecommitRepositoryInvalidDescriptionRule{
		resourceType:  "aws_codecommit_repository",
		attributeName: "description",
		max:           1000,
	}
}

// Name returns the rule name
func (r *AwsCodecommitRepositoryInvalidDescriptionRule) Name() string {
	return "aws_codecommit_repository_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodecommitRepositoryInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodecommitRepositoryInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodecommitRepositoryInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodecommitRepositoryInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1000 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
