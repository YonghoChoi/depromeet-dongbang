variable "cluster_name" {
  type = "string"
}

variable "project" {
  type = "string"
}

variable "environment" {
  type = "string"
}

variable "eks_bastion_instance_type" {
  type = "string"
}

variable "eks_node_instance_type" {
  type = "string"
}

variable "desired_capacity" {
  default = 1
}

variable "max_size" {
  default = 2
}

variable "min_size" {
  default = 1
}

variable "key_pair" {
  type = "string"
}

variable "cluster_version" {
  type = "string"
}

variable "volume_type" {
  type = "string"
  default = "gp2"
}

variable "volume_size" {
  type = "string"
  default = "50"
}

variable "ssh_password_parameter_name" {
  type = "string"
}

variable "aws_credential_parameter_name" {
  type = "string"
}

variable "eks_admin_arn" {
  type = "string"
}

variable "eks_admin_username" {
  type = "string"
}

variable "base_tags" {
  type = "map"
}

variable "eks_master_sg_ids" {
  type = "list"
}

variable "eks_node_sg_ids" {
  type = "list"
}

variable "eks_bastion_sg_ids" {
  type = "list"
}

variable "public_subnet_ids" {
  type = "list"
}

variable "private_subnet_ids" {
  type = "list"
}

variable "worker_node_labels" {
  type = "map"
  default = {
    "app": "default",
    "role": "worker",
  }
}