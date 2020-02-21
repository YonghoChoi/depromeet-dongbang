resource "aws_spot_instance_request" "bastion" {
  ami                  = data.aws_ami.ubuntu.id
  availability_zone    = data.aws_availability_zones.available.names[0]
  key_name             = var.key_pair
  instance_type        = var.eks_bastion_instance_type
  iam_instance_profile = aws_iam_instance_profile.eks_bastion.id
  vpc_security_group_ids = var.eks_bastion_sg_ids

  spot_price = "0.2"
  spot_type = "one-time"
  wait_for_fulfillment = "true"

  subnet_id                   = var.public_subnet_ids[0]
  associate_public_ip_address = true
  user_data = local.bastion_userdata

  tags = merge(
  map("key", "Name", "value", var.cluster_name),
  var.base_tags
  )

  connection {
    type        = "ssh"
    user        = "ubuntu"
    password    = data.aws_ssm_parameter.ec2_password.value
    host        = self.public_ip
  }

  provisioner "file" {
    source      = "${path.module}/../../../kubernetes"
    destination = "~/k8s"
  }

  provisioner "remote-exec" {
  inline = [
    "curl -o aws-iam-authenticator https://amazon-eks.s3-us-west-2.amazonaws.com/1.13.8/2019-08-14/bin/linux/amd64/aws-iam-authenticator",
    "mv ./aws-iam-authenticator /usr/bin/",
    "aws-iam-authenticator version",

    "mkdir -p ~/.kube",
    "sudo curl --silent --location -o /usr/local/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl",
    "sudo chmod +x /usr/local/bin/kubectl",
    "sudo kubectl version",
    "echo '${local.kubeconfig}' | tee ~/.kube/config &> /dev/null",
    "echo '${local.eks_configmap}' | tee ~/configmap.yaml &> /dev/null",
    "dos2unix ~/.kube/config",
    "dos2unix ~/configmap.yaml",
    "dos2unix ~/k8s/*.sh",
    "chmod +x ~/k8s/*.sh",

    "mkdir -p ~/.aws",
    "echo '${data.aws_ssm_parameter.aws_credential.value}' | tee ~/.aws/credentials &> /dev/null",

    "kubectl apply -f ~/configmap.yaml"
  ]
}

depends_on = ["aws_eks_cluster.cluster"]

timeouts {
  create = "30m"
  delete = "30m"
}
}

locals {
  bastion_userdata = <<USERDATA
#!/bin/bash

# docker install
apt-get update
apt-get install -y apt-transport-https ca-certificates curl software-properties-common awscli unzip dos2unix

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update
apt-get install -y docker-ce
usermod -aG docker ubuntu
curl -L "https://github.com/docker/compose/releases/download/1.23.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose

# set password
echo "ubuntu:${data.aws_ssm_parameter.ec2_password.value}" | chpasswd
sed -i "/^[^#]*PasswordAuthentication[[:space:]]no/c\PasswordAuthentication yes" /etc/ssh/sshd_config
service sshd restart
USERDATA
}