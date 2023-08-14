variable "tag_bucket_environment" {
  type        = string
  description = "Environment tag"
  default     = "Dev"
}

variable "tag_bucket_name" {
  type        = string
  description = "Bucket name"
  default     = "mytestbucket-05082023"
}

variable "s3_media" {
  type        = string
  description = "Audio"
  default     = "Type of media stored in S3 bucket"
}