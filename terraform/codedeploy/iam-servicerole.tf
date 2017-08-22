resource "aws_iam_role_policy" "codedeploy-mongosnap" {
  name = "codedeploy-mongosnap"
  role = "${aws_iam_role.codedeploy-mongosnap.id}"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
          "Action": [
                "autoscaling:CompleteLifecycleAction",
                "autoscaling:DeleteLifecycleHook",
                "autoscaling:DescribeAutoScalingGroups",
                "autoscaling:DescribeLifecycleHooks",
                "autoscaling:PutLifecycleHook",
                "autoscaling:RecordLifecycleActionHeartbeat",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "tag:GetTags",
                "tag:GetResources",
                "sns:Publish",
                "s3:GetObject",
                "s3:ListObject"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role" "codedeploy-mongosnap" {
  name = "codedeploy-mongosnap"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "codedeploy.amazonaws.com"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}
