// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule checks the pattern is valid
type AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule returns new rule with default attributes
func NewAwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule() *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule {
	return &AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule{
		resourceType:  "aws_launch_template",
		attributeName: "instance_initiated_shutdown_behavior",
		enum: []string{
			"stop",
			"terminate",
		},
	}
}

// Name returns the rule name
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Name() string {
	return "aws_launch_template_invalid_instance_initiated_shutdown_behavior"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLaunchTemplateInvalidInstanceInitiatedShutdownBehaviorRule) Check(runner *tflint.Runner) error {
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
					`instance_initiated_shutdown_behavior is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
