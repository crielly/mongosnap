provider "aws" {
  profile = "tastycidr"
}

variable "region" {
  default = "us-east-1"
}

variable "mongosnap-artifact-bucket" {}

variable "codedeploy-slack-endpoint" {}
