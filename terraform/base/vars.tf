provider "aws" {
  profile = "${var.aws-credentials-profile}"
}

variable "region" {
  default = "us-east-1"
}

variable "aws-credentials-profile" {}

variable "artifact-bucket" {}

variable "state-bucket" {}

variable "github-oauth-token" {}

data "terraform_remote_state" "codebuild" {
  backend = "s3"

  config {
    bucket  = "${var.state-bucket}"
    key     = "dev/codebuild.tfstate"
    profile = "${var.aws-credentials-profile}"
  }
}

data "terraform_remote_state" "codedeploy" {
  backend = "s3"

  config {
    bucket  = "${var.state-bucket}"
    key     = "dev/codedeploy.tfstate"
    profile = "${var.aws-credentials-profile}"
  }
}

data "terraform_remote_state" "codepipeline" {
  backend = "s3"

  config {
    bucket  = "${var.state-bucket}"
    key     = "dev/codepipeline.tfstate"
    profile = "${var.aws-credentials-profile}"
  }
}
