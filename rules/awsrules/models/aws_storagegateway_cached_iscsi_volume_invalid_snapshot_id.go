// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule checks the pattern is valid
type AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule returns new rule with default attributes
func NewAwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule() *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule {
	return &AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule{
		resourceType:  "aws_storagegateway_cached_iscsi_volume",
		attributeName: "snapshot_id",
		pattern:       regexp.MustCompile(`^\Asnap-([0-9A-Fa-f]{8}|[0-9A-Fa-f]{17})\z$`),
	}
}

// Name returns the rule name
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule) Name() string {
	return "aws_storagegateway_cached_iscsi_volume_invalid_snapshot_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsStoragegatewayCachedIscsiVolumeInvalidSnapshotIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`snapshot_id does not match valid pattern ^\Asnap-([0-9A-Fa-f]{8}|[0-9A-Fa-f]{17})\z$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
