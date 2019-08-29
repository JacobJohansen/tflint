// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule() *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule {
	return &AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "cname_prefix",
		max:           63,
		min:           4,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_cname_prefix"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidCnamePrefixRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"cname_prefix must be 63 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"cname_prefix must be 4 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
