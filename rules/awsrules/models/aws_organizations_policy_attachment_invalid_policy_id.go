// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule checks the pattern is valid
type AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsOrganizationsPolicyAttachmentInvalidPolicyIDRule returns new rule with default attributes
func NewAwsOrganizationsPolicyAttachmentInvalidPolicyIDRule() *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule {
	return &AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule{
		resourceType:  "aws_organizations_policy_attachment",
		attributeName: "policy_id",
		pattern:       regexp.MustCompile(`^p-[0-9a-zA-Z_]{8,128}$`),
	}
}

// Name returns the rule name
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Name() string {
	return "aws_organizations_policy_attachment_invalid_policy_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOrganizationsPolicyAttachmentInvalidPolicyIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`policy_id does not match valid pattern ^p-[0-9a-zA-Z_]{8,128}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
