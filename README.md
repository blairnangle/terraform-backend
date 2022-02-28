# terraform-backend

Provisions a backend store for [remote Terraform state](https://www.terraform.io/docs/state/remote.html) using AWS S3. 
State file locking is achieved using a DynamoDB table.

Creates one S3 bucket with a subdirectory for each Terraform project that references that bucket as its remote state store.

### AWS credentials

The file located at `~/.aws/credentials` should take the form:

```bash
[personal]
region=eu-west-2
aws_access_key_id=AKXXXXXXXXXXXXXXXXXX
aws_secret_access_key=3mFD7tVeyIakOOGH8A0v0rtZ/a3M+CyAs3mifl19
aws_secret_access_key=YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY
```

## Usage

The resources in this project should only need to be created once per AWS account. The remote backend created by this 
project will be used by this project as well as other Terraform projects. However, it cannot be referenced by this 
project before it has been created, so the Terraform commands are wrapped in a `go` script.

Set the `AWS_PROFILE` to the environment in which the resources are to be created and execute `go` from the relevant 
subdirectory. For the `personal` AWS account, for example:

```bash
$ export AWS_PROFILE=personal
```

```bash
$ ./go
```

The backend for the project is now S3. State files (`terraform.tfstate` and `terraform.tfstate.backup`) should 
have been synced to the remote backend and their local copies can be removed as well as the temporary plan, `plan.out`.

### Updating the Backend

Use the regular `terraform` commands (`plan` and `apply`) from the relevant environment subdirectory.

### In Other Projects

To make use of the backend store in another project named `blairnangle-dot-com`, for example, add the following to its 
Terraform configuration:

```hcl-terraform
terraform {
  backend "s3" {
    bucket         = "terraform-state-blair-nangle"
    key            = "blairnangle-dot-com/terraform.tfstate"
    region         = "eu-west-2"
    dynamodb_table = "terraform-locks"
    encrypt        = true
  }
}
```

## Gotchas

* `terraform init` will use the `default` profile in `~/.aws/credentials` (if set) unless the `AWS_PROFILE` environment 
variable has been set
* `terraform init` can only be run from the current directory and will parse all files with a `.tf` filetype, so the 
`backend` configuration needs to be created on the fly the very first time `terraform init` is executed (it cannot 
exist before the S3 bucket and DynamoDB table have been created and cannot be configured by passing command line 
arguments to `terraform init`)
* The resources associated with this project cannot be deleted using `terraform destroy` because S3 buckets cannot be 
deleted while they have contents and Terraform needs a state file so that it knows what needs to be deleted, so there is 
a circular dependency
