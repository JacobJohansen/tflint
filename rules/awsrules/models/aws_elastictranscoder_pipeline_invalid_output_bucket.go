// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsElastictranscoderPipelineInvalidOutputBucketRule checks the pattern is valid
type AwsElastictranscoderPipelineInvalidOutputBucketRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastictranscoderPipelineInvalidOutputBucketRule returns new rule with default attributes
func NewAwsElastictranscoderPipelineInvalidOutputBucketRule() *AwsElastictranscoderPipelineInvalidOutputBucketRule {
	return &AwsElastictranscoderPipelineInvalidOutputBucketRule{
		resourceType:  "aws_elastictranscoder_pipeline",
		attributeName: "output_bucket",
		pattern:       regexp.MustCompile(`^(\w|\.|-){1,255}$`),
	}
}

// Name returns the rule name
func (r *AwsElastictranscoderPipelineInvalidOutputBucketRule) Name() string {
	return "aws_elastictranscoder_pipeline_invalid_output_bucket"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastictranscoderPipelineInvalidOutputBucketRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastictranscoderPipelineInvalidOutputBucketRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastictranscoderPipelineInvalidOutputBucketRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastictranscoderPipelineInvalidOutputBucketRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`output_bucket does not match valid pattern ^(\w|\.|-){1,255}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
