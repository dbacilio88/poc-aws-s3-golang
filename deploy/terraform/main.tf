locals {
  object_source = "${path.module}/files/"
}

terraform {
  required_version = ">= 1.0.0" # Ensure that the Terraform version is 1.0.0 or higher
  required_providers {
    aws = {
      source = "hashicorp/aws" # Specify the source of the AWS provider
      version = "~> 4.0"        # Use a version of the AWS provider that is compatible with version
    }
  }
}

provider "aws" {
  region = var.region_id # Set the AWS region to US East (N. Virginia)
  access_key = var.user_ak
  secret_key = var.user_sk
}
