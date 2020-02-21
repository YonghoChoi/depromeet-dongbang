provider "aws" {
  region = "ap-northeast-2"
  profile = "dongbang_eksadmin"  // EKS 클러스터 생성자 = kubernetes admin
}

// backend 설정에 변수 지정 불가능
terraform {
  backend "s3" {
    region = "ap-northeast-2"
    bucket = "dongbang"
    key = "live/tfstates/ap-northeast-2.tfstate"
    encrypt = false
  }
}

module "dongbang" {
  source = "../modules/eks"
  cluster_version = "1.14"
  key_pair = "yongho42natecom"
  project = local.project
  environment = local.env
  cluster_name = local.cluster_name
  eks_admin_arn = "arn:aws:iam::992189553983:user/depromeet-dongbang"
  eks_admin_username = "depromeet-dongbang"
  eks_bastion_sg_ids = [module.eks-bastion.this_security_group_id]
  eks_master_sg_ids = [module.eks-master.this_security_group_id]
  eks_node_sg_ids = [module.eks-node.this_security_group_id]
  private_subnet_ids = module.vpc.private_subnets
  public_subnet_ids = module.vpc.public_subnets
  eks_bastion_instance_type = "c4.large"
  eks_node_instance_type = "c4.large"
  ssh_password_parameter_name = local.ssh_password_parameter_name
  aws_credential_parameter_name = "dongbang-credential"
  base_tags = local.base_tags
}


resource "null_resource" "wait_for_cluster" {
  depends_on = [
    module.dongbang
  ]

  connection {
    type        = "ssh"
    user        = "ubuntu"
    password    = data.aws_ssm_parameter.ec2_password.value
    host        = module.dongbang.bastion_public_ip
  }

  provisioner "remote-exec" {
    inline = [
      "mkdir ~/mongodata",
      "cd ~/k8s",
      "./create.sh"
    ]
  }
}
