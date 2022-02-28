terraform {
  backend "s3" {
    bucket         = "terraform-state-blair-nangle"
    key            = "terraform-backend/terraform.tfstate"
    region         = "eu-west-2"
    dynamodb_table = "terraform-locks"
    encrypt        = true
  }
}
