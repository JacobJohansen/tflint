// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsLambdaLayerVersionInvalidS3ObjectVersionRule checks the pattern is valid
type AwsLambdaLayerVersionInvalidS3ObjectVersionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsLambdaLayerVersionInvalidS3ObjectVersionRule returns new rule with default attributes
func NewAwsLambdaLayerVersionInvalidS3ObjectVersionRule() *AwsLambdaLayerVersionInvalidS3ObjectVersionRule {
	return &AwsLambdaLayerVersionInvalidS3ObjectVersionRule{
		resourceType:  "aws_lambda_layer_version",
		attributeName: "s3_object_version",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionInvalidS3ObjectVersionRule) Name() string {
	return "aws_lambda_layer_version_invalid_s3_object_version"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionInvalidS3ObjectVersionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaLayerVersionInvalidS3ObjectVersionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionInvalidS3ObjectVersionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionInvalidS3ObjectVersionRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"s3_object_version must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"s3_object_version must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
