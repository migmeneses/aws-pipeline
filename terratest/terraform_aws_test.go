package test

import (
        "testing"

        "github.com/gruntwork-io/terratest/modules/aws"
        "github.com/gruntwork-io/terratest/modules/terraform"
        "github.com/stretchr/testify/assert"
)

func TestTerraformAwsFlugel(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Reviewing EC2 and S3 tags
	// Run `terraform output` to get the value of an output variable
        instanceID := terraform.Output(t, terraformOptions, "instance_id")
	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	// Look up the tags for the given Instance and Bucket IDs
	instanceTags := aws.GetTagsForEc2Instance(t, "eu-central-1", instanceID)
        bucketTags := aws.GetS3BucketTags(t, "eu-central-1", bucketID)

	// Verify that our expected name tag is one of the tags
        ec2NameTag, ec2ContainsNameTag := instanceTags["Name"]
        assert.True(t, ec2ContainsNameTag)
	assert.Equal(t, "Flugel", ec2NameTag)
	s3NameTag, s3ContainsNameTag := bucketTags["Name"]
        assert.True(t, s3ContainsNameTag)
        assert.Equal(t, "Flugel", s3NameTag)

        // Verify that our expected owner tag is one of the tags
	ec2OwnerTag, ec2ContainsOwnerTag := instanceTags["Owner"]
        assert.True(t, ec2ContainsOwnerTag)
        assert.Equal(t, "InfraTeam", ec2OwnerTag)
        s3OwnerTag, s3ContainsOwnerTag := bucketTags["Owner"]
        assert.True(t, s3ContainsOwnerTag)
        assert.Equal(t, "InfraTeam", s3OwnerTag)
}
