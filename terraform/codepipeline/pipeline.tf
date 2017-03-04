# resource "aws_codepipeline" "mongosnap" {
#   name     = "mongosnap"
#   role_arn = "${aws_iam_role.mongosnap-codepipeline.arn}"
#   artifact_store {
#     location = "${aws_s3_bucket.mongosnap.bucket}"
#     type     = "S3"
#   }
#   stage {
#     name = "Source"
#     action {
#       name             = "Source"
#       category         = "Source"
#       owner            = "crielly"
#       provider         = "GitHub"
#       version          = "1"
#       output_artifacts = ["mongosnap"]
#       configuration {
#         Owner  = "crielly"
#         Repo   = "https://github.com/crielly/mongosnap.git"
#         Branch = "master"
#       }
#     }
#   }
#   stage {
#     name = "Build"
#     action {
#       name            = "Build"
#       category        = "Build"
#       owner           = "AWS"
#       provider        = "CodeBuild"
#       input_artifacts = ["mongosnap"]
#       version         = "1"
#       configuration {
#         ProjectName = "test"
#       }
#     }
#   }
# }

