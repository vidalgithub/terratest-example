package test

import (
	"crypto/tls"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestS3Website(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	websiteEndpoint := "http://" + terraform.Output(t, terraformOptions, "website_endpoint")

	tlsConfig := tls.Config{}

	maxRetries := 5
	timeBetweenRetries := 5 * time.Second

	instanceText := "<H1>Hello World!</H1>"

	http_helper.HttpGetWithRetry(t, websiteEndpoint, &tlsConfig, 200, instanceText, maxRetries, timeBetweenRetries)

}
