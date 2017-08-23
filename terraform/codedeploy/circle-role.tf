resource "aws_iam_policy" "circle" {
  name        = "${terraform.workspace}-circleci"
  description = "${terraform.workspace}-circleci"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:*"
            ],
            "Resource": [
                "${aws_s3_bucket.mongosnap.arn}/*",
                "${aws_s3_bucket.mongosnap.arn}/"
            ]
        }
    ]
}
EOF
}
