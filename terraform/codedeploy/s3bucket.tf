resource "aws_s3_bucket" "mongosnap" {
    bucket = "${var.NAMESPACE}-mongosnap-${var.REGION}"

    tags {
        Name        = "${var.NAMESPACE}-mongodump-${var.REGION}"
        Environment = "core"
        Terraform   = "True"
    }
}
