// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsElasticBeanstalkApplicationVersionInvalidApplicationRule checks the pattern is valid
type AwsElasticBeanstalkApplicationVersionInvalidApplicationRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkApplicationVersionInvalidApplicationRule returns new rule with default attributes
func NewAwsElasticBeanstalkApplicationVersionInvalidApplicationRule() *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule {
	return &AwsElasticBeanstalkApplicationVersionInvalidApplicationRule{
		resourceType:  "aws_elastic_beanstalk_application_version",
		attributeName: "application",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule) Name() string {
	return "aws_elastic_beanstalk_application_version_invalid_application"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkApplicationVersionInvalidApplicationRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"application must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"application must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
