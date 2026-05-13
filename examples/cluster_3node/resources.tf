# =============================================================================
# Three-Node NetScaler Cluster
# =============================================================================
#
# This example creates an L2 cluster with three NetScaler ADC nodes using
# the composite citrixadc_cluster resource.
#
# The resource handles:
#   1. Bootstrapping the first node (node_id=0) as the Cluster Coordinator (CCO)
#   2. Creating the Cluster IP (CLIP)
#   3. Joining the remaining nodes to the cluster
#
# Node with the lowest priority value becomes the CCO.
# Priority is computed as 30 + node_id (i.e., 30, 31, 32).
#
# Usage:
#   terraform init
#   terraform apply
# =============================================================================

locals {
  cluster_nodes = [
    for name, attrs in var.netscaler_attributes : {
      name      = name
      nodeid    = attrs.node_id
      ip        = var.nsips[name]
      priority  = 29 + attrs.node_id
      backplane = attrs.backplane
    }
  ]

  # Sort by node_id to ensure deterministic ordering (CCO first)
  cluster_nodes_sorted = [
    for s in sort([
      for n in local.cluster_nodes : format("%03d|%s", n.nodeid, jsonencode(n))
    ]) : jsondecode(split("|", s)[1])
  ]
}

resource "citrixadc_cluster" "this" {
  clid          = var.cluster_id
  clip          = var.clip
  hellointerval = 200

  # Allow extra time for each node to stabilize before adding the next
  bootstrap_poll_delay    = "60s"
  bootstrap_poll_interval = "60s"
  bootstrap_total_timeout = "10m"
  node_add_poll_delay     = "30s"
  node_add_poll_interval  = "30s"
  node_add_total_timeout  = "10m"

  dynamic "clusternode" {
    for_each = local.cluster_nodes_sorted
    content {
      nodeid     = clusternode.value.nodeid
      endpoint   = "http://${clusternode.value.ip}"
      ipaddress  = clusternode.value.ip
      backplane  = clusternode.value.backplane
      priority   = clusternode.value.priority
      state      = "ACTIVE"
      tunnelmode = "NONE"
    }
  }
}
