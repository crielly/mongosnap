resource "aws_security_group" "mongosnap" {
  name                = "${terraform.workspace}-mongosnap"
  description         = "${terraform.workspace}-mongosnap"

  ingress {
    from_port       = 27017
    to_port         = 27017
    protocol        = "tcp"
    self            = true
    cidr_blocks = ["${var.VPNCIDR}"]
  }

  ingress {
    from_port       = 22
    to_port         = 22
    protocol        = "tcp"
    cidr_blocks = ["${var.VPNCIDR}"]
  }

  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  } 

  tags {
    Name            = "${terraform.workspace}-mongosnap"
    Environment     = "${terraform.workspace}"
    Terraform       = "True"
  }
}
