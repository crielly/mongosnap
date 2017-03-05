resource "aws_codepipeline" "mongosnap" {
  name     = "mongosnap"
  role_arn = "${aws_iam_role.mongosnap-codepipeline.arn}"

  artifact_store {
    location = "${aws_s3_bucket.mongosnap.bucket}"
    type     = "S3"
  }

  stage {
    name = "Source"

    action {
      name             = "github-checkout"
      category         = "Source"
      owner            = "ThirdParty"
      provider         = "GitHub"
      version          = "1"
      output_artifacts = ["mongosnap-source"]

      configuration {
        OAuthToken = "${var.github-oauth-token}"
        Owner      = "crielly"
        Repo       = "mongosnap"
        Branch     = "master"
      }
    }
  }

  stage {
    name = "Build"

    action {
      name             = "mongosnap-compile"
      category         = "Build"
      owner            = "AWS"
      provider         = "CodeBuild"
      input_artifacts  = ["mongosnap-source"]
      output_artifacts = ["mongosnap-build"]
      version          = "1"

      configuration {
        ProjectName = "${data.terraform_remote_state.codebuild.project-name}"
      }
    }
  }

  stage {
    name = "Deploy"

    action {
      name            = "binary-deploy"
      category        = "Deploy"
      owner           = "AWS"
      provider        = "CodeDeploy"
      input_artifacts = ["mongosnap-build"]
      version         = "1"

      configuration {
        ApplicationName     = "${data.terraform_remote_state.codedeploy.application-name}"
        DeploymentGroupName = "${data.terraform_remote_state.codedeploy.deployment-group-name}"
      }
    }
  }
}
