resource "aws_codedeploy_app" "mongosnap" {
  name = "mongosnap"
}

resource "aws_codedeploy_deployment_config" "mongosnap" {
  deployment_config_name = "mongosnap"

  minimum_healthy_hosts {
    type  = "HOST_COUNT"
    value = 1
  }
}

resource "aws_codedeploy_deployment_group" "mongosnap-dev" {
  app_name               = "${aws_codedeploy_app.mongosnap.name}"
  deployment_group_name  = "mongosnap-dev"
  service_role_arn       = "${aws_iam_role.mongosnap-codedeploy.arn}"
  deployment_config_name = "${aws_codedeploy_deployment_config.mongosnap.id}"

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
