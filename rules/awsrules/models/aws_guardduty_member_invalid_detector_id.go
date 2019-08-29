// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsGuarddutyMemberInvalidDetectorIDRule checks the pattern is valid
type AwsGuarddutyMemberInvalidDetectorIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyMemberInvalidDetectorIDRule returns new rule with default attributes
func NewAwsGuarddutyMemberInvalidDetectorIDRule() *AwsGuarddutyMemberInvalidDetectorIDRule {
	return &AwsGuarddutyMemberInvalidDetectorIDRule{
		resourceType:  "aws_guardduty_member",
		attributeName: "detector_id",
		max:           300,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyMemberInvalidDetectorIDRule) Name() string {
	return "aws_guardduty_member_invalid_detector_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyMemberInvalidDetectorIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyMemberInvalidDetectorIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyMemberInvalidDetectorIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyMemberInvalidDetectorIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"detector_id must be 300 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"detector_id must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
