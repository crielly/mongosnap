resource "aws_instance" "mongosnap" {
  ami                     = "${data.aws_ami.latest-hvm-1404-ami.id}"
  ebs_optimized           = false
  instance_type           = "t2.micro"
  monitoring              = false
  key_name                = "${aws_key_pair.provisioner.key_name}"
  iam_instance_profile    = "${aws_iam_instance_profile.mongosnap.name}"
  disable_api_termination = true
  source_dest_check       = "false"

  vpc_security_group_ids = [
    "${aws_security_group.mongosnap.id}"
  ]

  associate_public_ip_address = true

  root_block_device {
    volume_type           = "standard"
    volume_size           = 8
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "/dev/xvdb"
    volume_type           = "gp2"
    volume_size           = 5
    delete_on_termination = true
  }

  ebs_block_device {
    device_name           = "/dev/xvdc"
    volume_type           = "gp2"
    volume_size           = 5
    delete_on_termination = true
  }

  tags {
    Environment    = "${terraform.workspace}"
    Name           = "${terraform.workspace}-mongosnap"
    Terraform      = "True"
    mongosnap      = "True"
  }

  lifecycle {
    ignore_changes = ["ami"]
  }
}
