resource "aws_codebuild_project" "mongosnap" {
  name         = "mongosnap"
  description  = "builds mongosnap golang binary"
  timeout      = "30"
  service_role = "${aws_iam_role.codebuild_role.arn}"

  artifacts {
    type = "CODEPIPELINE"
  }

  environment {
    compute_type = "BUILD_GENERAL1_SMALL"
    image        = "2"
    type         = "LINUX_CONTAINER"
  }

  source {
    type = "CODEPIPELINE"
  }

  tags {
    "Environment" = "dev"
  }
}
