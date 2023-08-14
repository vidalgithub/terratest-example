package test

import (
	"encoding/json"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3IsVersioned(t *testing.T) {
	awsRegion := "eu-central-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")

	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)

}

func TestGetS3BucketTagsV1(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	tagsMap := terraform.OutputMap(t, terraformOptions, "tags")

	tagsToCheck := []string{"Environment", "Name"}

	filteredTags := make(map[string]string)

	for _, tag := range tagsToCheck {
		if value, ok := tagsMap[tag]; ok {
			filteredTags[tag] = value
		}
	}

	expectedTagsString := `{"Environment":"Dev","Name":"mytestbucket-05082023"}`

	var expectedTags map[string]string
	err := json.Unmarshal([]byte(expectedTagsString), &expectedTags)
	if err != nil {
		t.Fatalf("Failed to unmarshal expected tags: %s", err)
	}

	assert.Equal(t, expectedTags, filteredTags)

}
