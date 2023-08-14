resource "aws_s3_bucket" "test_bucket" {
  bucket = "mytestbucket-05082023"

  tags = {
    Name        = var.tag_bucket_name
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
