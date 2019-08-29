package awsrules

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/terraform/configs"
	"github.com/hashicorp/terraform/configs/configload"
	"github.com/hashicorp/terraform/terraform"
	"github.com/wata727/tflint/client"
	"github.com/wata727/tflint/tflint"
)

func Test_AwsALBInvalidSubnet(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Response []*ec2.Subnet
		Expected tflint.Issues
	}{
		{
			Name: "subnet ID is invalid",
			Content: `
resource "aws_alb" "balancer" {
    subnets = [
        "subnet-1234abcd",
        "subnet-abcd1234",
    ]
}`,
			Response: []*ec2.Subnet{
				{
					SubnetId: aws.String("subnet-12345678"),
				},
				{
					SubnetId: aws.String("subnet-abcdefgh"),
				},
			},
			Expected: tflint.Issues{
				{
					Rule:    NewAwsALBInvalidSubnetRule(),
					Message: "\"subnet-1234abcd\" is invalid subnet ID.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 4, Column: 9},
						End:      hcl.Pos{Line: 4, Column: 26},
					},
				},
				{
					Rule:    NewAwsALBInvalidSubnetRule(),
					Message: "\"subnet-abcd1234\" is invalid subnet ID.",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 9},
						End:      hcl.Pos{Line: 5, Column: 26},
					},
				},
			},
		},
		{
			Name: "subnet ID is valid",
			Content: `
resource "aws_alb" "balancer" {
    subnets = [
        "subnet-1234abcd",
        "subnet-abcd1234",
    ]
}`,
			Response: []*ec2.Subnet{
				{
					SubnetId: aws.String("subnet-1234abcd"),
				},
				{
					SubnetId: aws.String("subnet-abcd1234"),
				},
			},
			Expected: tflint.Issues{},
		},
		{
			Name: "use list variables",
			Content: `
variable "subnets" {
    default = ["subnet-1234abcd", "subnet-abcd1234"]
}

resource "aws_alb" "balancer" {
    subnets = "${var.subnets}"
}`,
			Response: []*ec2.Subnet{
				{
					SubnetId: aws.String("subnet-1234abcd"),
				},
				{
					SubnetId: aws.String("subnet-abcd1234"),
				},
			},
			Expected: tflint.Issues{},
		},
	}

	dir, err := ioutil.TempDir("", "AwsALBInvalidSubnet")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(currentDir)

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for _, tc := range cases {
		loader, err := configload.NewLoader(&configload.Config{})
		if err != nil {
			t.Fatal(err)
		}

		err = ioutil.WriteFile(dir+"/resource.tf", []byte(tc.Content), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}

		mod, diags := loader.Parser().LoadConfigDir(".")
		if diags.HasErrors() {
			t.Fatal(diags)
		}
		cfg, tfdiags := configs.BuildConfig(mod, configs.DisabledModuleWalker)
		if tfdiags.HasErrors() {
			t.Fatal(tfdiags)
		}

		runner, err := tflint.NewRunner(tflint.EmptyConfig(), map[string]tflint.Annotations{}, cfg, map[string]*terraform.InputValue{})
		if err != nil {
			t.Fatal(err)
		}
		rule := NewAwsALBInvalidSubnetRule()

		mock := client.NewMockEC2API(ctrl)
		mock.EXPECT().DescribeSubnets(&ec2.DescribeSubnetsInput{}).Return(&ec2.DescribeSubnetsOutput{
			Subnets: tc.Response,
		}, nil)
		runner.AwsClient.EC2 = mock

		if err = rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		opts := []cmp.Option{
			cmpopts.IgnoreUnexported(AwsALBInvalidSubnetRule{}),
			cmpopts.IgnoreFields(hcl.Pos{}, "Byte"),
		}
		if !cmp.Equal(tc.Expected, runner.Issues, opts...) {
			t.Fatalf("Expected issues are not matched:\n %s\n", cmp.Diff(tc.Expected, runner.Issues, opts...))
		}
	}
}
