resource "aws_s3_bucket" "mongosnap" {
  bucket = "${var.artifact-bucket}"

  tags {
    Name = "mongosnap"
  }
}
