output "bastion_public_ip" {
  value = aws_spot_instance_request.bastion.public_ip
}