resource "aws_eks_node_group" "worker" {
  cluster_name = var.cluster_name
  node_group_name = "${var.cluster_name}-worker"
  node_role_arn = aws_iam_role.eks_node.arn
  subnet_ids = var.private_subnet_ids
  instance_types = [
    var.eks_node_instance_type]
  disk_size = var.volume_size
  labels = var.worker_node_labels

  tags = merge(
  map("key", "Name", "value", "${var.cluster_name}-worker"),
  var.base_tags
  )

  remote_access {
    ec2_ssh_key = var.key_pair
    source_security_group_ids = var.eks_node_sg_ids
  }

  scaling_config {
    desired_size = var.desired_capacity
    max_size = var.max_size
    min_size = var.min_size
  }

  # Ensure that IAM Role permissions are created before and deleted after EKS Node Group handling.
  # Otherwise, EKS will not be able to properly delete EC2 Instances and Elastic Network Interfaces.
  depends_on = [
    aws_iam_role_policy_attachment.eks_worker_node,
    aws_iam_role_policy_attachment.eks_cni,
    aws_iam_role_policy_attachment.ecr_ro,
  ]
}