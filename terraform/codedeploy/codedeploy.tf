resource "aws_codedeploy_app" "mongosnap" {
  name = "mongosnap"
}

resource "aws_codedeploy_deployment_group" "mongosnap" {
  app_name               = "${aws_codedeploy_app.mongosnap.name}"
  deployment_group_name  = "${terraform.workspace}-mongosnap"
  service_role_arn       = "${aws_iam_role.codedeploy-mongosnap.arn}"
  deployment_config_name = "CodeDeployDefault.AllAtOnce"

  ec2_tag_filter {
    key   = "mongosnap"
    type  = "KEY_AND_VALUE"
    value = "True"
  }

  ec2_tag_filter {
    key   = "environment"
    type  = "KEY_AND_VALUE"
    value = "${terraform.workspace}"
  }

  auto_rollback_configuration {
    enabled = true
    events  = ["DEPLOYMENT_FAILURE"]
  }
}

output "application-name" {
  value = "${aws_codedeploy_app.mongosnap.name}"
}

output "deployment-group-name" {
  value = "${aws_codedeploy_deployment_group.mongosnap.deployment_group_name}"
}
