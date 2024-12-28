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

resource "aws_s3_bucket" "s3_golang" {
  bucket = var.name_bucket
  tags = {
    Name = "my bucket"
    Environment = "Dev"
    Language = "Golang"
  }
}

resource "aws_s3_object" "s3_golang_files" {
  for_each = fileset(local.object_source, "*")
  bucket = aws_s3_bucket.s3_golang.id
  key    = each.value
  source = "${local.object_source}/${each.value}"
  etag = filemd5("${local.object_source}/${each.value}")
}