terraform {
  backend "s3" {
    bucket = "mongosnap-tfstate"
    key    = "backend/codepipeline.tfstate"
    region = "us-east-1"
  }
}
