resource "aws_codedeploy_app" "mongosnap" {
  name = "mongosnap"
}

resource "aws_codedeploy_deployment_group" "mongosnap-dev" {
  app_name               = "${aws_codedeploy_app.mongosnap.name}"
  deployment_group_name  = "mongosnap-dev"
  service_role_arn       = "${aws_iam_role.codedeploy-role.arn}"
  deployment_config_name = "CodeDeployDefault.OneAtATime"

  ec2_tag_filter {
    key   = "mongosnap"
    type  = "KEY_AND_VALUE"
    value = "true"
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
  value = "${aws_codedeploy_deployment_group.mongosnap-dev.deployment_group_name}"
}
