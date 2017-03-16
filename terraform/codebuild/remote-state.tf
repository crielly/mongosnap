terraform {
  backend "s3" {
    bucket = "mongosnap-tfstate"
    key    = "backend/codebuild.tfstate"
    region = "us-east-1"
  }
}
