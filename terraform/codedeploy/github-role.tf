resource "aws_iam_role_policy" "github-mongosnap-codedeploy" {
  name = "github-mongosnap-codedeploy"
  role = "${aws_iam_role.github-mongosnap-codedeploy.id}"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
          "codedeploy:RegisterApplicationRevision",
          "codedeploy:GetApplicationRevision",
          "codedeploy:CreateDeployment"
        ],
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role" "github-mongosnap-codedeploy" {
  name = "github-mongosnap-codedeploy"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "AWS": [
          "arn:aws:iam::458265684782:user/codedeploy"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}