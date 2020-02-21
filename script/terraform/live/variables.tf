locals {
  project = "dongbang"
  env = "live"
  cluster_name = "${local.project}-${local.env}"
  base_tags = {
    Project = local.project
    Environment = local.env
    Terraform = "true"
  }
  ssh_password_parameter_name = "dongbang-password"
}

