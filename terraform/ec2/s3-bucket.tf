resource "aws_s3_bucket" "mongodump" {
  bucket = "${var.NAMESPACE}-mongodump-${var.REGION}"

  lifecycle_rule {
    id      = "daily"
    prefix  = "Daily/"
    enabled = true

    expiration {
      days = 7
    }
  }

  lifecycle_rule {
    id      = "weekly"
    prefix  = "Weekly/"
    enabled = true

    expiration {
      days = 7
    }
  }

  lifecycle_rule {
    id      = "monthly"
    prefix  = "Monthly/"
    enabled = true

    transition {
      days          = 2
      storage_class = "GLACIER"
    }

    expiration {
      days = 7
    }
  }

  tags {
    Name        = "${var.NAMESPACE}-mongodump-${var.REGION}"
    Environment = "core"
    Terraform   = "True"
  }
}
