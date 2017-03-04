provider "aws" {
  profile = "tastycidr"
}

variable "region" {
  default = "us-east-1"
}

variable "mongosnap-artifact-bucket" {}

variable "codedeploy-slack-endpoint" {}

data "terraform_remote_state" "codebuild" {
  backend = "s3"

  config {
    bucket  = "mongosnap-tfstate"
    key     = "dev/codebuild.tfstate"
    profile = "tastycidr"
  }
}

data "terraform_remote_state" "codedeploy" {
  backend = "s3"

  config {
    bucket  = "mongosnap-tfstate"
    key     = "dev/codedeploy.tfstate"
    profile = "tastycidr"
  }
}

data "terraform_remote_state" "codepipeline" {
  backend = "s3"

  config {
    bucket  = "mongosnap-tfstate"
    key     = "dev/codepipeline.tfstate"
    profile = "tastycidr"
  }
}
