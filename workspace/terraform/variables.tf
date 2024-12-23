variable "name_bucket" {
  type = string
  default = "my-s3-golang-bucket"
  description = "Name bucket"
  nullable = false
  sensitive = true
}