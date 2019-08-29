// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsBackupVaultInvalidNameRule checks the pattern is valid
type AwsBackupVaultInvalidNameRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsBackupVaultInvalidNameRule returns new rule with default attributes
func NewAwsBackupVaultInvalidNameRule() *AwsBackupVaultInvalidNameRule {
	return &AwsBackupVaultInvalidNameRule{
		resourceType:  "aws_backup_vault",
		attributeName: "name",
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9\-\_\.]{1,50}$`),
	}
}

// Name returns the rule name
func (r *AwsBackupVaultInvalidNameRule) Name() string {
	return "aws_backup_vault_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBackupVaultInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBackupVaultInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBackupVaultInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBackupVaultInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-zA-Z0-9\-\_\.]{1,50}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
