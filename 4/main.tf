resource "aws_s3_bucket" "static_website" {
  bucket = "mytestbucket-05082023"

  tags = {
    Name        = var.tag_bucket_name
    Environment = var.tag_bucket_environment
    Media = var.s3_media
  }
}

resource "aws_s3_bucket_public_access_block" "my_website" {
  bucket = aws_s3_bucket.static_website.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

resource "aws_s3_bucket_website_configuration" "online" {
  bucket                = aws_s3_bucket.static_website.bucket
  expected_bucket_owner = "532199187081"

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }

}

resource "aws_s3_object" "index_html" {
  bucket = aws_s3_bucket.static_website.bucket
  key    = "index.html"
  source       = "./website/index.html"
  content_type = "text/html"
}

resource "aws_s3_object" "error_html" {
  bucket = aws_s3_bucket.static_website.bucket
  key    = "error.html"
  source       = "./website/error.html"
  content_type = "text/html"
}

resource "aws_s3_bucket_policy" "allow_access" {
  bucket = aws_s3_bucket.static_website.id
  policy = data.aws_iam_policy_document.allow_access.json

  depends_on = [ 
    aws_s3_bucket.static_website, 
    data.aws_iam_policy_document.allow_access,
    aws_s3_bucket_public_access_block.my_website
  ]
}

resource "aws_s3_bucket_versioning" "test_bucket" {
  bucket = aws_s3_bucket.static_website.id
  versioning_configuration {
    status = "Enabled"
  }
}

data "aws_iam_policy_document" "allow_access" {
  statement {
    sid = "PublicReadGetObject"

    principals {
      type        = "AWS"
      identifiers = ["*"]
    }

    actions = [
      "s3:GetObject",
      "s3:ListBucket",
    ]

    resources = [
      aws_s3_bucket.static_website.arn,
      "${aws_s3_bucket.static_website.arn}/*",
    ]
  }
}

output "website_endpoint" {
  value = aws_s3_bucket_website_configuration.online.website_endpoint
}

output "bucket_id" {
  value = aws_s3_bucket.static_website.id
}

output "tags" {
  value = aws_s3_bucket.static_website.tags
}