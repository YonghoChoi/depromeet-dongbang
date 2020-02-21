data "aws_availability_zones" "available" {
  state = "available"
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"
  version = "2.17.0"

  name = local.cluster_name
  cidr = "10.0.0.0/16"

  azs = data.aws_availability_zones.available.names
  private_subnets = [
    "10.0.32.0/19",
    "10.0.64.0/19",
    "10.0.96.0/19"]
  public_subnets = [
    "10.0.128.0/19",
    "10.0.160.0/19",
    "10.0.192.0/19"]

  enable_dns_support = true
  enable_dns_hostnames = true
  single_nat_gateway = true
  enable_nat_gateway = true

  public_subnet_tags = {
    "kubernetes.io/role/elb" = "1",
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
  }

  private_subnet_tags = {
    "kubernetes.io/role/internal-elb" = "1"
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
  }

  tags = merge(
    local.base_tags,
    {
      "Name" = local.cluster_name
    },
  )
}

module "eks-master" {
  source = "terraform-aws-modules/security-group/aws"

  name = "eks-master-${local.cluster_name}-sg"
  description = "EKS master security group"
  vpc_id = module.vpc.vpc_id

  ingress_with_source_security_group_id = [
    {
      rule = "https-443-tcp"
      source_security_group_id = module.eks-bastion.this_security_group_id
    },
    {
      rule = "https-443-tcp"
      source_security_group_id = module.eks-node.this_security_group_id
    }
  ]

  ingress_with_cidr_blocks = [
    {
      rule = "https-443-tcp"
      description = "my local ip"
      cidr_blocks = "${chomp(data.http.myip.body)}/32"
    }
  ]

  egress_with_cidr_blocks = [
    {
      rule = "all-all"
      cidr_blocks = "0.0.0.0/0"
    }
  ]

  tags = merge(
    local.base_tags,
    {
      "Name" = "eks-master-${local.cluster_name}-sg",
      "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    },
  )
}

module "eks-node" {
  source = "terraform-aws-modules/security-group/aws"

  name = "eks-node-${local.cluster_name}-sg"
  description = "EKS node security group"
  vpc_id = module.vpc.vpc_id

  ingress_with_cidr_blocks = [
    {
      from_port = 30000
      to_port = 32767
      protocol = "tcp"
      description = "vpc cidr (node port range)"
      cidr_blocks = module.vpc.vpc_cidr_block
    }
  ]

  ingress_with_source_security_group_id = [
    {
      rule = "all-all"
      source_security_group_id = module.eks-bastion.this_security_group_id
    },
    {
      rule = "all-all"
      source_security_group_id = module.eks-node.this_security_group_id
    },
    {
      from_port = 1025
      to_port = 65535
      protocol = "tcp"
      description = "master node sg"
      source_security_group_id = module.eks-master.this_security_group_id
    },
    {
      from_port = 443
      to_port = 443
      protocol = "tcp"
      description = "master node sg"
      source_security_group_id = module.eks-master.this_security_group_id
    },
    {
      from_port = 30000
      to_port = 32767
      protocol = "tcp"
      description = "bastion server sg (node port range)"
      source_security_group_id = module.eks-bastion.this_security_group_id
    }
  ]

  egress_with_cidr_blocks = [
    {
      rule = "all-all"
      cidr_blocks = "0.0.0.0/0"
    }
  ]

  tags = merge(
    local.base_tags,
    {
      "Name" = "eks-node-${local.cluster_name}-sg",
      "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    },
  )
}

module "eks-bastion" {
  source = "terraform-aws-modules/security-group/aws"

  name = "eks-bastion-${local.cluster_name}-sg"
  description = "EKS bastion security group"
  vpc_id = module.vpc.vpc_id

  ingress_with_cidr_blocks = [
    {
      rule = "ssh-tcp"
      description = "my local ip"
      cidr_blocks = "${chomp(data.http.myip.body)}/32"
    }
  ]

  egress_with_cidr_blocks = [
    {
      rule = "all-all"
      cidr_blocks = "0.0.0.0/0"
    }
  ]

  tags = merge(
    local.base_tags,
    {
      "Name" = "eks-bastion-${local.cluster_name}-sg",
    },
  )
}
