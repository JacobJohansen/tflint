// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsMacieS3BucketAssociationInvalidMemberAccountIDRule checks the pattern is valid
type AwsMacieS3BucketAssociationInvalidMemberAccountIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsMacieS3BucketAssociationInvalidMemberAccountIDRule returns new rule with default attributes
func NewAwsMacieS3BucketAssociationInvalidMemberAccountIDRule() *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule {
	return &AwsMacieS3BucketAssociationInvalidMemberAccountIDRule{
		resourceType:  "aws_macie_s3_bucket_association",
		attributeName: "member_account_id",
		pattern:       regexp.MustCompile(`^[0-9]{12}$`),
	}
}

// Name returns the rule name
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Name() string {
	return "aws_macie_s3_bucket_association_invalid_member_account_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsMacieS3BucketAssociationInvalidMemberAccountIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`member_account_id does not match valid pattern ^[0-9]{12}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
