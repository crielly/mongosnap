resource "aws_s3_bucket" "mongosnap" {
  bucket = "${var.mongosnap-artifact-bucket}"

  tags {
    Name = "mongosnap"
  }
}
