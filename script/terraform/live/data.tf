data "http" "myip" {
  url = "http://ipv4.icanhazip.com"
}

data "aws_ssm_parameter" "ec2_password" {
  name = local.ssh_password_parameter_name
}