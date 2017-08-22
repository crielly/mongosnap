resource "aws_iam_role" "mongosnap" {
  name = "${terraform.workspace}-mongosnap"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "mongosnap" {
  name        = "${terraform.workspace}-mongosnap"
  description = "${terraform.workspace}-mongosnap"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:Get*",
                "s3:List*",
                "s3:ListBucket"
            ],
            "Resource": [
                "${aws_s3_bucket.mongodump.arn}/*",
                "${aws_s3_bucket.mongodump.arn}/"
            ]
        }
    ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "mongosnap" {
  role       = "${aws_iam_role.mongosnap.name}"
  policy_arn = "${aws_iam_policy.mongosnap.arn}"
}

resource "aws_iam_instance_profile" "mongosnap" {
  name = "${terraform.workspace}-mongosnap"
  role = "${aws_iam_role.mongosnap.name}"
}
