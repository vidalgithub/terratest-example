package test

import (
	"crypto/tls"
	"testing"
	"time"

	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/aws/aws-sdk-go/service/s3/s3manager"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestS3Website(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Get the tags output as a map[string]string.
	websiteEndpoint := "http://" + terraform.Output(t, terraformOptions, "website_endpoint")

	// Setup a TLS configuration to submit with the helper, a blank struct is acceptable
	tlsConfig := tls.Config{}

	// It can take a minute or so for the Instance to boot up, so retry a few times
	maxRetries := 2
	timeBetweenRetries := 5 * time.Second

	// Expected values
	instanceText := "<H1>Hello World!</H1>"

	// Verify that we get back a 200 OK with the expected instanceText
	http_helper.HttpGetWithRetry(t, websiteEndpoint, &tlsConfig, 200, instanceText, maxRetries, timeBetweenRetries)

}
