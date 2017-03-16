terraform {
  backend "s3" {
    bucket = "mongosnap-tfstate"
    key    = "backend/codedeploy.tfstate"
    region = "us-east-1"
  }
}
