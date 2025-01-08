variable "name_bucket" {
  type = string
  default = "my-s3-golang-bucket"
  description = "Name bucket"
  nullable = false
  sensitive = true
}

variable "region_id" {
  type = string
  default = ""
  description = "region id aws s3"
  nullable = false
  sensitive = true
}

variable "user_ak" {
  type = string
  default = ""
  description = "access key aws s3"
  nullable = false
  sensitive = true
}

variable "user_sk" {
  type = string
  default = ""
  description = "secret key aws s3"
  nullable = false
  sensitive = true
}

variable "name_dynamodb" {
type = string
  default = ""
  description = "Name database"
  nullable = false
  sensitive = true
}