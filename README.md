Solving the DevOps Infrastructure Dilemma: Enabling developer velocity with control 💡

Register for the webinar here →

How it works

Solutions
Documentation
Pricing

Resources
Login
Book a demo
Start for Free
TERRAFORM
What is Terratest and How to Use it

What is Terratest and How to Use it
Sumeet Ninawe
03 Oct 2023
·
20 min read
Reviewed by: 
Flavius Dinu
terratest
In this post, we explore various concepts around Terratest and how to use Terratest in various testing scenarios. In case you are not familiar with Go programming language, I have tried my best to explain the code in this post.

We will cover:

What is Terratest?
Static vs. dynamic testing
How does Terratest work?
Writing unit tests with Terratest
Writing integration tests with Terratest
Terratest stages
Terratest vs. Terragrunt
Terratest alternatives
Note: All the source code is found in this monorepo. Link to the appropriate folder will be highlighted in relevant sections.

What is Terratest?
Terratest is an open source testing framework for infrastructure defined as code using Terraform. It performs unit tests, integration tests, and end-to-end tests for the cloud-based infrastructure and helps identify security vulnerabilities early on. It is possible to automate and integrate Terratest with Terraform CI/CD workflow, making it convenient for developers to receive early feedback during development.

Terratest features
Some of the features of Terratest are described below.

Infrastructure testing automation: Terratest automates testing for Terraform, enabling efficient and consistent validation of cloud resources and configurations.
Multi-cloud support: It offers compatibility with various cloud providers, including AWS, Azure, and Google Cloud, allowing users to test across different platforms.
Programmatic test definition: Developers can write tests in Go programming language, enabling expressive and code-based test cases that interact with Terraform-managed resources.
Testing levels: Terratest supports a range of testing levels, from unit and integration tests to end-to-end scenarios, ensuring comprehensive coverage of infrastructure code behavior.
Early issue detection: By catching problems before deployment, Terratest enhances the reliability of infrastructure code, reducing risks and promoting more stable cloud environments.
Static vs. dynamic testing
Before we proceed, let us understand the difference between static and dynamic testing in the context of IaC.

Earlier, we covered the details about tfsec, which is also known as a static testing tool for Terraform IaC. Tfsec scans for Terraform configuration files in the given directory and identifies security shortcomings based on the community-defined and custom-defined test cases locally. It does not require an internet connection since it does not actually provision the cloud infrastructure in the real world.

On the other hand, Terratest is dynamic because it actually provisions the live infrastructure and helps run custom tests defined using SDK in Go programming language. The infrastructure thus created is also cleaned up in the same test run. Naturally, executing Terratest takes longer since time is spent on creating the infrastructure.

Both tfsec and Terratest – static and dynamic – have their virtues. Tfsec is a great tool to leverage the community-driven (as well as custom-defined) best practices in IaC practices, thus giving an offset while addressing infrastructure security issues. While Terratest helps test the real infrastructure based on a test library provided in the SDK.

To learn more about testing Terraform code check out our How to Test Terraform Code article and see other useful Terraform tools.

How does Terratest work?
As mentioned earlier, Terratest test cases are defined using the Go programming language. Specifically the testing package of Go. This testing package is used by Go developers to define test cases in general.

The Terratest Go library leverages this to define a variety of test cases used to test Docker images, cloud infrastructure defined for AWS, Azure, GCP, Kubernetes, and many more.

1. Define a basic Terraform configuration file
To introduce the basic working of Terratest, let us define a basic Terraform configuration file that just outputs a string value, as shown below.

output "tftest_output" {
  value = "Hello Terratest!"
}
2. Run Terraform apply
Running terraform apply simply prints the value above.

% terraform apply

Changes to Outputs:
  + tftest_output = "Hello Terratest!"

You can apply this plan to save these new output values to the Terraform state, without changing any
real infrastructure.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes


Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

Outputs:

tftest_output = "Hello Terratest!"
3. Define a test case using Terratest
To define a test case using Terratest, create a “tests” directory in the project root. This is for better file organization.

You can also include all the Terraform config files in a separate directory, but for the examples discussed here, we only create a separate directory for Go test files. Within the tests directory, create a file named tf_test.go. 

It is important to end the filename with “_test.go” – this restriction is imposed by the Go testing package. It helps Go to identify the files containing the code for tests.

The current directory structure looks like shown below. We would follow the same structure for all other examples in this post.

terratest gruntwork
As seen in the screenshot above (directory 1), find the code and configuration files at this location.

The tf_test.go file contains the code below. We will go through this code step by step, and try to understand the rudiments of defining Terratest test cases.

package test

import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformHelloWorldExample(t *testing.T) {
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../",
    })

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)

    output := terraform.Output(t, terraformOptions, "tftest_output")
    assert.Equal(t, "Hello Terratest!", output)
}
Every Go code should be packaged. The first line in the code above defines the package name. This package name can be anything. Here we have defined the package as “test”.
Next we import a few packages. As mentioned before, Terratest depends on the Go testing package, so we have imported the same. We have also imported the Terrratest’s Terraform module, which helps us run various stages of each test. Finally, we have imported a 3rd party library to use its “assert” function. The assert function is generally used in testing frameworks to compare the expected and actual outputs.
The function named “TestTerraformHelloWorldExample” is where our test function definition begins. The name of the function has to begin with “Test” for Go to execute it as part of the testing package. This function accepts a pointer to testing.T struct – this goes into the details of how Go structs work. For now, just understand that this object is required to execute various test functions.
The terraformOptions within the “TestTerraformHelloWorldExample” function is initialized with terraform.WithDefaultRetryableErrors() function. This is used to initialize several parameters required by Terraform. In this example, we have provided the path to the directory where Terraform config files reside. All the terraform operations like init, apply, destroy, work based on the terraformOptions variable.
terraform.Destroy() function – as the name suggests – destroys all the infrastructure created by executing this code. The defer keyword is a Go specific feature that defers the execution of this function until the surrounding code/functions are executed.
terraform.InitAndApply() function initializes the Terraform project and applies the same to provision the real world cloud infrastructure resources. In this case, we have only defined an output variable. So there is no real infrastructure creation, but we use the value of this output to test our results.
The terraform.Output() function reads the values of the output variables defined in the Terraform configuration. In our case, we have defined the “tftest_output” output variable and hardcoded its value to “Hello Terratest!”. This is the value using which the output variable is initialized.
Finally, we compare the output variable with the assigned value with our expected value. The expected value is passed as the 2nd parameter to assert.Equal() function in the last line of code. Optionally, it is also possible to assign this value to a variable named “expectedOutput” for better readability.
In general, the values to be tested are made available to Terratest test cases via Terraform’s output variable. These values are then compared (asserted) against the expected values, and the result of the test case is determined.

4. Run “go test” observe the output
To execute this test case, navigate to the tests directory run go test and observe the output.

The steps performed after running the go test are as follows:

Terraform project is initialized.
Terraform configuration is applied, and output values are obtained.
The obtained output values are then compared with the expected results.
The infrastructure is destroyed.
Test results are presented.
.
.
.
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: Changes to Outputs:
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66:   - tftest_output = "Hello Terratest!" -> null
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: 
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: You can apply this plan to save these new output values to the Terraform
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: state, without changing any real infrastructure.
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: 
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: Destroy complete! Resources: 0 destroyed.
TestTerraformHelloWorldExample 2023-08-14T15:09:00+02:00 logger.go:66: 
PASS
ok      terratest-example/tests 3.911s
As seen above, the Test result is “PASS,” and it also shows which tests were performed and how long it took to perform the test.

By now, this should have given you a better understanding of how to implement Terratest test cases.

Writing unit tests with Terratest - Example
Now that we are familiar with the process of writing tests in Terratest, in this section, we create real infrastructure components and execute a couple of test cases.

The Terraform config below creates an S3 bucket with the given name and tags. It also enables versioning on this bucket and outputs bucket ID and tags information via output variables.

resource "aws_s3_bucket" "test_bucket" {
  bucket = "mytestbucket-05082023"

  tags = {
    Name = var.tag_bucket_name
    Environment = var.tag_bucket_environment
    Media = var.s3_media
  }
}

resource "aws_s3_bucket_versioning" "test_bucket" {
  bucket = aws_s3_bucket.test_bucket.id
  versioning_configuration {
    status = "Enabled"
  }
}

output bucket_id {
  value = aws_s3_bucket.test_bucket.id
}

output tags {
  value = aws_s3_bucket.test_bucket.tags
}
Here, we have to test a couple of things:

Check if the versioning is enabled on the bucket
Check for specific tag values being set. We want to make sure that the below tags are set amongst others:
Environment: Dev
Name: mytestbucket-05082023
The function TestS3IsVersioned() below defines the steps to check for versioning on the S3 bucket.

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
Most of the steps in the function above are similar to the Terratest test defined in the previous example. However, there are a few differences. We explicitly get the versioning information using Golang AWS SDK using the bucket ID. The bucket ID is exposed by the Terraform config as part of its configuration. This becomes the actual value of for assertion.

If the versioning is enabled on the S3 bucket defined in the Terraform configuration, the test passes.

The second test is defined in the TestGetS3BucketTagsV1() function. Here, the aim is to compare the tags returned by the Terraform output variable “tags”, with expected values stored in the “expectedTagsString” variable in Golang code below.

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
A bucket may have multiple tags defined. The scope of our test is limited to the “Environment” and “Name” tags only. Thus the output variable “tags” may return more than required tags.

The Go code implements this filter starting from the line where tagsToCheck is initialized, till the end of the for loop.

The Terraform configuration defines an additional tag named “Media”, which is not of our interest. So these lines of code pick only the Environment and Name tag and store it in a struct named filteredTags.

The rest of the code in this function performs a direct comparison of the filteredTags with expectedTags.

When the above test is run using the go test command, it reports the following output in the terminal.

.
.
.
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: aws_s3_bucket_versioning.test_bucket: Destruction complete after 0s
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: aws_s3_bucket.test_bucket: Destroying... [id=mytestbucket-05082023]
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: aws_s3_bucket.test_bucket: Destruction complete after 1s
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: 
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: Destroy complete! Resources: 2 destroyed.
TestGetS3BucketTagsV1 2023-08-14T15:23:28+02:00 logger.go:66: 
PASS
ok      tests3  36.660s
Here we have successfully developed unit test cases for the given Terraform configuration.

Notice that this time the test run took 36.6 seconds to complete. This is significantly higher than the previous test. Most of this time is consumed in the provisioning and destruction of the S3 bucket in AWS.

Also note that since we have defined two different functions, the bucket is provisioned and destroyed twice. This is indeed a bit of an overhead, and we will address the same in upcoming sections.

The code for this example is found here.

Learn how to manage S3 buckets with Terraform.

💡 You might also like:

5 Ways to Manage Terraform at Scale
How to Improve Your Infrastructure as Code using Terraform
How to Automate Terraform Deployments and Infrastructure Provisioning
Writing integration tests with Terratest - Example
Integration tests are a type of testing that assesses the interactions between different components or modules within a system. These tests aim to uncover issues that might arise when multiple components are deployed and are expected to interact, ensuring the smooth functionality of the entire system. Integration tests help detect integration-related bugs, data flow problems, and communication issues early in the development process.

The steps to write integration tests using Terratest are similar to what we have been doing till now. However, in this section, we will demonstrate how various features of Terratest can be leveraged to implement custom and integration testing.

In our example, the Terraform configuration for the S3 bucket is updated to enable static website hosting on the same. We want to write a test that makes sure that the static web hosting is indeed enabled and the test website is accessible from the internet. To do this, we use the “http-helper” module of Terratest to make a GET request to the website, which is hosted based on the given Terraform config.

You can find the complete updated Terraform config here.

There are many solutions available on the internet for hosting a Static website using Terraform on AWS S3. For the sake of this example, we will focus on the output variables exposed by the Terraform config below.

output "website_endpoint" {
  value = aws_s3_bucket_website_configuration.online.website_endpoint
}

output "bucket_id" {
  value = aws_s3_bucket.static_website.id
}

output "tags" {
  value = aws_s3_bucket.static_website.tags
}
This Terraform configuration now exposes an additional variable named “website_endpoint”. We will use this endpoint to send a GET request in our test.

In the tf_test.go file, we have defined a function to perform this particular integration test, as seen below.

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
The first few steps till terraform.InitAndApply() are similar to previous examples – they create the S3 bucket, enable static website hosting, and upload index.html and error.html files to this bucket. Once the tests are over, the S3 bucket is destroyed.

We use the output variable to construct an endpoint URL in the form of a string and initialize the websiteEndpoint variable.

Terratest’s http_helper.HttpGetWithRetry() function performs the test of comparing the HTML contents returned from the endpoint URL, with the expected string stored in the instanceText variable.

This function also takes additional noteworthy parameters like 

maxRetries – number of trials performed to get a success result before the test is failed. This is important on several occasions as some cloud resources may take longer to be provisioned.
timeBetweenRetries – wait time between trials.
Run go test and observe the output. In the output logs on the terminal, we can see that:

Terraform is being initialized.
S3 bucket is created, and the corresponding static website configuration is enabled.
The output variables are printed, and the endpoint URL is constructed.
GET call is made to the constructed endpoint.
The response is compared, and subsequently, the bucket is destroyed.
Test results are printed.
.
.
.
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: Outputs:
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: 
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: bucket_id = "mytestbucket-05082023"
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: tags = tomap({
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66:   "Environment" = "Dev"
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66:   "Media" = "Type of media stored in S3 bucket"
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66:   "Name" = "mytestbucket-05082023"
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: })
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: website_endpoint = "mytestbucket-05082023.s3-website.eu-central-1.amazonaws.com"
TestS3Website 2023-08-14T19:45:07+02:00 retry.go:91: terraform [output -no-color -json website_endpoint]
TestS3Website 2023-08-14T19:45:07+02:00 logger.go:66: Running command terraform with args [output -no-color -json website_endpoint]
TestS3Website 2023-08-14T19:45:08+02:00 logger.go:66: "mytestbucket-05082023.s3-website.eu-central-1.amazonaws.com"
TestS3Website 2023-08-14T19:45:08+02:00 retry.go:91: HTTP GET to URL http://mytestbucket-05082023.s3-website.eu-central-1.amazonaws.com
TestS3Website 2023-08-14T19:45:08+02:00 http_helper.go:59: Making an HTTP GET call to URL http://mytestbucket-05082023.s3-website.eu-central-1.amazonaws.com
TestS3Website 2023-08-14T19:45:08+02:00 retry.go:91: terraform [destroy -auto-approve -input=false -lock=false]
.
.
.
TestS3Website 2023-08-14T19:45:17+02:00 logger.go:66: aws_s3_bucket.static_website: Destroying... [id=mytestbucket-05082023]
TestS3Website 2023-08-14T19:45:17+02:00 logger.go:66: aws_s3_bucket.static_website: Destruction complete after 1s
TestS3Website 2023-08-14T19:45:17+02:00 logger.go:66: 
TestS3Website 2023-08-14T19:45:17+02:00 logger.go:66: Destroy complete! Resources: 7 destroyed.
TestS3Website 2023-08-14T19:45:17+02:00 logger.go:66: 
PASS
ok      tests3  21.367s
Terratest stages
Until now, we have dealt with relatively small Terraform configurations. In the real world, Terraform projects can grow and the infrastructure components being deployed would be large in numbers.

In such cases, writing Terratest tests in a similar manner is not a good idea for several reasons.

Infrastructure components often are dependent on other components. Integration tests may fail unnecessarily if all the required components are not ready.
Parallel execution can cause resource contention.
Typically integration tests are run after unit tests are successful. In our examples, we have implemented them separately. It should be possible to define a sequence for such tests.
As seen in the second example, for two test functions defined, the infrastructure is created twice, which is not the best way to write tests.
To address the above concerns, in this section, we introduce the test_structure feature of Terratest.

Using test_structure, we can define the sequence of tests and perform all the tests without provisioning and destroying the infrastructure multiple times.

Note that this is just the beginning of writing effective tests using Terratest. There are many more features offered by Terratest that are leveraged to write better tests with minimum overhead.

The Terraform config used in this example is the same as that of the previous section. However, this time we will combine all the tests together to demonstrate test_structure feature.

The complete code is available here.

We will perform the unit tests to check for versioning and tags, and then we also perform the integration test to confirm if the static website is up using the S3 bucket.

The overall test plan is represented in the diagram below.

Terratest stages
The TestS3 function acts as a “main” function. It provisions and destroys the infrastructure and also implements test_structure to define the sequence of tests to be performed.

As per this sequence, 

bucketVersionValidation() function validates if the versioning is enabled on the S3 bucket.
tagsValidation() function checks for the specific tags and the corresponding expected values.
endpointValidation() function performs the integration test by checking if the URL returns the expected HTML string.
The validation steps required for each test are wrapped in separate functions to improve readability. The output produced by Terraform configuration in the TestS3 function is passed to appropriate functions.

The TestS3 function looks like below.

func TestS3(t *testing.T) {
    now := time.Now()
    expectedName := fmt.Sprintf("mytestbucket-%s", strings.ToLower(now.Format("01022006")))

    expectedEnvironment := "Dev"

    awsRegion := "eu-central-1"

    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        TerraformDir: "../",

        Vars: map[string]interface{}{
        "tag_bucket_name": expectedName,
        "tag_bucket_environment": expectedEnvironment,
        },
    })

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)

    // Check versioning
    test_structure.RunTestStage(t, "test_versioning", func() {
        bucketID := terraform.Output(t, terraformOptions, "bucket_id")
        bucketVersionValidation(t, terraformOptions, awsRegion, bucketID)
    })

    // Tags comparison
    test_structure.RunTestStage(t, "test_tags", func() {
        tagsValidation(t, terraformOptions)
    })

    // Endpoint testing
    test_structure.RunTestStage(t, "endpoint_test", func() {
        endpointValidation(t, terraformOptions)
    })
}
terraformOptions is initialized with an additional attribute “Vars” which will be used by some of the test functions.
The test_structure.RunTestStage() function helps run the tests in stages/sequence. The name of each stage is specified in the 2nd argument passed in each call. Finally, this function accepts an anonymous function, which inturn calls the corresponding function to execute steps.
Notice how the relevant parameters are passed in the relevant functions and also how the functions being called from test_structure.RunTestStage() do not begin with “Test”.

The code snippet below shows these functions. The explanations are already covered in the previous sections. Here we are just splitting the logic into separate fragments.

func bucketVersionValidation(t *testing.T, terraformOptions *terraform.Options, awsRegion string, bucketID string) {
    // Bucket versioning comparison
    actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
    expectedStatus := "Enabled"
    assert.Equal(t, expectedStatus, actualStatus)
}

func tagsValidation(t *testing.T, terraformOptions *terraform.Options) {
    tagsMap := terraform.OutputMap(t, terraformOptions, "tags")

    tagsToCheck := []string{"Environment", "Name"}

    filteredTags := make(map[string]string)

    for _, tag := range tagsToCheck {
        if value, ok := tagsMap[tag]; ok {
             filteredTags[tag] = value
        }
    }

expectedTagsString := `{"Environment":"Dev","Name":"mytestbucket-08142023"}`

    var expectedTags map[string]string
    err := json.Unmarshal([]byte(expectedTagsString), &expectedTags)
    if err != nil {
        t.Fatalf("Failed to unmarshal expected tags: %s", err)
    }

    assert.Equal(t, expectedTags, filteredTags)
}

func endpointValidation(t *testing.T, terraformOptions *terraform.Options) {
    websiteEndpoint := "http://" + terraform.Output(t, terraformOptions, "website_endpoint")

    tlsConfig := tls.Config{}

    maxRetries := 2
    timeBetweenRetries := 5 * time.Second

    instanceText := "<H1>Hello World!</H1>"

    http_helper.HttpGetWithRetry(t, websiteEndpoint, &tlsConfig, 200, instanceText, maxRetries, timeBetweenRetries)

}
Terratest vs. Terragrunt
Terratest and Terragrunt are important tools in the realm of infrastructure as code (IaC), specifically Terraform, for managing and deploying infrastructure on cloud platforms.

Terratest focuses on automated testing of Terraform configurations, enabling developers to validate their infrastructure code by writing tests that simulate real-world scenarios. It helps catch issues early in the development process, ensuring the reliability and stability of infrastructure changes. 

Terragrunt addresses the complexities of managing multiple Terraform deployments, offering a higher level of abstraction for managing configurations and remote state. It aids in maintaining a consistent and organized infrastructure by allowing the separation of code into reusable modules and enforcing best practices across projects.

While Terratest streamlines testing, Terragrunt simplifies the management of complex infrastructures, making both tools invaluable for enhancing the efficiency and quality of infrastructure management workflows.

Check out also how Spacelift makes it easy to work with Terraform and Terragrunt. If you need any help managing your Terraform infrastructure, building more complex workflows based on Terraform, and managing AWS credentials per run, instead of using a static pair on your local machine, Spacelift is a fantastic tool for this.

It supports Git workflows, policy as code, programmatic configuration, context sharing, drift detection, and many more great features right out of the box. See how you can integrate security tools using custom inputs, and if you want to learn more about Spacelift, create a free account today or book a demo with one of our engineers.

Terratest alternatives
Below is the list of Terratest alternatives in the space of IaC and not limited to Terraform.

Kitchen-Terraform: Integrates Terraform testing into the popular Test Kitchen framework, facilitating multi-platform testing of infrastructure code.
ServerSpec: A Ruby-based testing framework that verifies infrastructure state and configuration on remote servers.
Molecule: Primarily used for testing Ansible roles, Molecule supports testing infrastructure code in different environments using multiple virtualization platforms.
Pulumi Test: Part of the Pulumi framework, it enables testing of infrastructure code written in Pulumi across different cloud providers.
Terraform Compliance: Focuses on security and compliance by allowing the creation of tests to validate whether Terraform configurations adhere to specific security standards.
InSpec: A tool for creating infrastructure tests that assess security compliance, availability, and other attributes of systems.
TerratestCDK: Extends Terratest to work with AWS Cloud Development Kit (CDK), combining CDK’s infrastructure modeling with Terratest’s testing capabilities.
Bats: A Bash Automated Testing System, suitable for simple infrastructure testing within shell scripts.
Goss: Specializes in validating server configuration through YAML or JSON tests, ensuring system state matches the desired setup.
Testinfra: Designed for infrastructure testing, it offers a Python-based framework to validate server properties and configurations.
Key points
In this post, we introduced Terratest and attempted to provide you with a way to create test cases for infrastructure defined using Terraform IaC. Terratest is a great tool that offers dynamic test capabilities for Terraform configuration. By extending Go’s testing package, it is possible to leverage the Terratest library to define highly customized tests. 

We also saw how multiple tests can be combined into a single test plan by leveraging the test_structure capabilities provided by Terratest. It should be noted that Terratest goes much beyond the features defined in this blog post. It is important to be aware of these capabilities, and their Github repo is a great place to start.

 

Terraform Management Made Easy
Spacelift effectively manages Terraform state, more complex workflows, supports policy as code, programmatic configuration, context sharing, drift detection, resource visualization and includes many more features.

Start free trial
Written by

avatar_image_sumeetn
Sumeet Ninawe
Sumeet has over ten years of overall experience in IT and has worked with cloud and DevOps technologies for the last four years. He is a Certified System Administrator and TOGAF® 9. He specializes in writing IaC using Terraform. In his free time, Sumeet maintains a blog at LetsDoTech.
Read also
OpenTofu: The Open-Source Alternative to Terraform
OPENTOFU
16 min read
OpenTofu: The Open-Source Alternative to Terraform
Terragrunt vs. Terraform &#8211; Comparison &#038; When to Use
TERRAFORM
12 min read
Terragrunt vs. Terraform – Comparison & When to Use
Managing Infrastructure as Code (IaC) With Terraform
TERRAFORM
16 min read
Managing Infrastructure as Code (IaC) With Terraform
Get our newsletter
Product
Documentation
How it works
Spacelift Tutorial
Pricing
Customer Case Studies
Integrations
Security
System Status
Product Updates
Test Pilot Program
Company
About Us
Careers
Contact Sales
Partners
Media resources
Learn
Blog
Atlantis Alternative
Terraform Cloud Alternative
Terraform Enterprise Alternative
Spacelift for AWS
Terraform Automation
© 2024 Spacelift, Inc. All rights reserved

Privacy PolicyTerms of Service
Struggling to balance developer velocity with control?

Attend the June 25 webinar:

Solving the DevOps Infrastructure Dilemma

Register for the webinar

