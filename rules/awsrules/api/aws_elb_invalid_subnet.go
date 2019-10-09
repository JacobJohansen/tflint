// This file generated by `tools/api-rule-gen/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/wata727/tflint/tflint"
)

// AwsELBInvalidSubnetRule checks whether attribute value actually exists
type AwsELBInvalidSubnetRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsELBInvalidSubnetRule returns new rule with default attributes
func NewAwsELBInvalidSubnetRule() *AwsELBInvalidSubnetRule {
	return &AwsELBInvalidSubnetRule{
		resourceType:  "aws_elb",
		attributeName: "subnets",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsELBInvalidSubnetRule) Name() string {
	return "aws_elb_invalid_subnet"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsELBInvalidSubnetRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsELBInvalidSubnetRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsELBInvalidSubnetRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeSubnets
func (r *AwsELBInvalidSubnetRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeSubnets")
			var err error
			r.data, err = runner.AwsClient.DescribeSubnets()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeSubnets",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		return runner.EachStringSliceExprs(attribute.Expr, func(val string, expr hcl.Expression) {
			if !r.data[val] {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is invalid subnet ID.`, val),
					expr.Range(),
				)
			}
		})
	})
}
