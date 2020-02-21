locals {
  project = "dongbang"
  env = "live"
  cluster_name = "${local.project}-${local.env}}"
  project = "dongbang"
  base_tags = {
    Project = local.project
    Environment = local.env
    Terraform = "true"
  }
}

