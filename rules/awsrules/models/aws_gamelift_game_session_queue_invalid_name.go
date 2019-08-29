// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsGameliftGameSessionQueueInvalidNameRule checks the pattern is valid
type AwsGameliftGameSessionQueueInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsGameliftGameSessionQueueInvalidNameRule returns new rule with default attributes
func NewAwsGameliftGameSessionQueueInvalidNameRule() *AwsGameliftGameSessionQueueInvalidNameRule {
	return &AwsGameliftGameSessionQueueInvalidNameRule{
		resourceType:  "aws_gamelift_game_session_queue",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[a-zA-Z0-9-]+$`),
	}
}

// Name returns the rule name
func (r *AwsGameliftGameSessionQueueInvalidNameRule) Name() string {
	return "aws_gamelift_game_session_queue_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGameliftGameSessionQueueInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGameliftGameSessionQueueInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGameliftGameSessionQueueInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGameliftGameSessionQueueInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[a-zA-Z0-9-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
