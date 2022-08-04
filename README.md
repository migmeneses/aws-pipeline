# TEST N.1

## Terraform Code
There is a terraform code that will deploy:
- 1x EC2 instance
- 1x S3 Bucket ( Simple Storage )

All of them get deployed in the region `eu-central-1`, which located in Frankfurt
Both resources have the tags: Name=Flugel, Owner=InfraTeam
And you can find the code in: `main.tf`


## Automatic test wirh Terratest
There is a tool called terratest, and on this occasion helps us to verify if the terraform output will give us the correct tagging \
for the resources mentioned above.

### Requirements
- An AWS account and AWS Access Credentials
- Terratest uses the Go testing framework. \
To use Terratest, you need to install: [Go](https://golang.org/)
- Terraform CLI. Ref - [Install Terraform](https://learn.hashicorp.com/tutorials/terraform/install-cli) 

### Running automated test
- Sign up for AWS.
- Configure your AWS credentials using one of the supported methods for AWS CLI tools. \
Example:  ~/.aws/config/credentials files contains `aws_access_key_id` and `aws_secret_access_key`
- Get to the right path and/or directory. On this occasion: flugel-it/terratest/
- Configure Go dependencies by running:
```bash
go mod init github.com/<YOUR_USERNAME>/<YOUR_REPO_NAME>
go mod tidy
```
- Run the Test ( this will run terraform init, apply, output and destroy and it will pass if we have the right tagging )
```bash
go test -v -run TestTerraformAwsFlugel 
```


## GitHub Actions Workflow
GitHub Actions provides us with a pipeline to verify if the code is correct and deploys infrastructure if been requested

### Requirements
- A Terraform Cloud account
- An AWS account and AWS Access Credentials

### API-TOKENS
We have 3x api-tokens in here:
- `AWS_ACCESS_KEY_ID` found `~/.aws/config/credentials`
- `AWS_SECRET_ACCESS_KEY` found `~/.aws/config/credentials`
- `TF_API_TOKEN`, it has to be created within the Terraform Cloud, under the User settings follow by [Tokens](https://app.terraform.io/app/settings/tokens?utm_source=learn)

### Secrets
On GitHub
- Go to the top right of the repo menu, and Select `Settings`
- On the left hand menu, please choose `Secrects`, followed by `Actions`
- Add 3x `New repository secret(s)`. `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` and `TF_API_TOKEN` with their respective api-tokens as values

### Run the Pipeline
Within the github repo, `.github/workflows/` you can find the `terraform.yml` file.
This file defines the workflow of the pipeline in steps.
Each step is running a different terraform command, and it ends by providing us with the infrastructure been requested

### How to run the pipeline
The pipeline runs automatically once you push into the main branch, or made a pull request
You can find the running pipeline under `Actions` on the top of the repo menu. Continue by:
- Click on the latest workflow run
- Click on Terraform, and see each job/step status

