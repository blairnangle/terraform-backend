#!/usr/bin/env bash

terraform init

terraform plan -out plan.out

terraform apply -auto-approve plan.out

echo "terraform {
  backend \"s3\" {
    bucket         = \"terraform-state-blair-nangle\"
    key            = \"terraform-backend/terraform.tfstate\"
    region         = \"eu-west-2\"
    dynamodb_table = \"terraform-locks\"
    encrypt        = true
  }
}" > backend.tf

terraform init -force-copy
