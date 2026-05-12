terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
    http = {
      source = "hashicorp/http"
    }
  }
}

# =============================================================================
# Two-Phase HA + LB Example
# =============================================================================
#
# This example demonstrates:
#   1. Setting up an HA pair between two NetScaler ADC nodes
#   2. Detecting which node became primary after HA formation
#   3. Configuring load balancing only on the primary node
#
# Usage (two-phase apply required):
#
#   # Phase 1: Create the HA pair
#   terraform apply \
#     -target=citrixadc_hanode.ha_peer_on_node1 \
#     -target=citrixadc_hanode.ha_peer_on_node2
#
#   # Phase 2: Detect primary and configure LB on it
#   terraform apply
#
# Why two phases?
#   Terraform evaluates data sources at plan time, so the HA pair must
#   already exist before we can read each node's HA state. The count
#   meta-argument requires a value known at plan time.
# =============================================================================

# -----------------------------------------------------------------------------
# Provider configuration – one alias per node
# -----------------------------------------------------------------------------

provider "citrixadc" {
  endpoint             = format("https://%s", var.node1_nsip)
  username             = var.adc_admin_username
  password             = var.adc_admin_password
  insecure_skip_verify = true
  alias                = "node1"
}

provider "citrixadc" {
  endpoint             = format("https://%s", var.node2_nsip)
  username             = var.adc_admin_username
  password             = var.adc_admin_password
  insecure_skip_verify = true
  alias                = "node2"
}

# =============================================================================
# PHASE 1 – HA Pair Formation
# =============================================================================
# Add each node as a peer of the other. The rpcnodepassword allows the nodes
# to authenticate over the HA heartbeat / sync channels.
# The node that initiates the add (node1) typically becomes Primary.
# -----------------------------------------------------------------------------

resource "citrixadc_hanode" "ha_peer_on_node1" {
  provider = citrixadc.node1

  hanode_id       = 1
  ipaddress       = var.node2_nsip
  rpcnodepassword = var.adc_admin_password
}

resource "citrixadc_hanode" "ha_peer_on_node2" {
  provider = citrixadc.node2

  hanode_id       = 1
  ipaddress       = var.node1_nsip
  rpcnodepassword = var.adc_admin_password

  depends_on = [citrixadc_hanode.ha_peer_on_node1]
}

# =============================================================================
# PHASE 2 – Primary Detection
# =============================================================================
# After the HA pair is formed (Phase 1), query each node's hanode/0 endpoint
# to discover which is Primary and which is Secondary.
# -----------------------------------------------------------------------------

data "http" "ha_node1" {
  url = format("http://%s/nitro/v1/config/hanode/0", var.node1_nsip)
  request_headers = {
    Authorization = "Basic ${base64encode("${var.adc_admin_username}:${var.adc_admin_password}")}"
    Accept        = "application/json"
  }
}

data "http" "ha_node2" {
  url = format("http://%s/nitro/v1/config/hanode/0", var.node2_nsip)
  request_headers = {
    Authorization = "Basic ${base64encode("${var.adc_admin_username}:${var.adc_admin_password}")}"
    Accept        = "application/json"
  }
}

locals {
  node1_ha = jsondecode(data.http.ha_node1.response_body)
  node2_ha = jsondecode(data.http.ha_node2.response_body)

  node1_is_primary = try(local.node1_ha.hanode[0].state == "Primary", false)
  node2_is_primary = try(local.node2_ha.hanode[0].state == "Primary", false)
}

output "node1_is_primary" {
  value = local.node1_is_primary
}

output "node2_is_primary" {
  value = local.node2_is_primary
}

output "primary_node_ip" {
  value = local.node1_is_primary ? var.node1_nsip : (local.node2_is_primary ? var.node2_nsip : "unknown")
}

# =============================================================================
# PHASE 2 – LB configuration on the primary node only
# =============================================================================
# In an HA pair, configuration applied to the primary auto-syncs to the
# secondary, so we only need to push config to whichever node is primary.
# -----------------------------------------------------------------------------

# --- Enable LB feature -------------------------------------------------------

resource "citrixadc_nsfeature" "lb_feature_node1" {
  provider = citrixadc.node1
  lb       = true
  count    = local.node1_is_primary ? 1 : 0
}

resource "citrixadc_nsfeature" "lb_feature_node2" {
  provider = citrixadc.node2
  lb       = true
  count    = local.node2_is_primary ? 1 : 0
}

# --- Backend servers ----------------------------------------------------------

resource "citrixadc_server" "backend_node1" {
  provider = citrixadc.node1
  count    = local.node1_is_primary ? length(var.backend_servers) : 0

  name      = var.backend_servers[count.index].name
  ipaddress = var.backend_servers[count.index].ip
}

resource "citrixadc_server" "backend_node2" {
  provider = citrixadc.node2
  count    = local.node2_is_primary ? length(var.backend_servers) : 0

  name      = var.backend_servers[count.index].name
  ipaddress = var.backend_servers[count.index].ip
}

# --- LB Virtual Server -------------------------------------------------------

resource "citrixadc_lbvserver" "app_lb_node1" {
  provider    = citrixadc.node1
  count       = local.node1_is_primary ? 1 : 0

  name        = var.lb_vserver_name
  ipv46       = var.lb_vip
  port        = var.lb_port
  servicetype = var.lb_service_type
  lbmethod    = var.lb_method
}

resource "citrixadc_lbvserver" "app_lb_node2" {
  provider    = citrixadc.node2
  count       = local.node2_is_primary ? 1 : 0

  name        = var.lb_vserver_name
  ipv46       = var.lb_vip
  port        = var.lb_port
  servicetype = var.lb_service_type
  lbmethod    = var.lb_method
}

# --- Services bound to the LB vserver ----------------------------------------

resource "citrixadc_service" "backend_svc_node1" {
  provider    = citrixadc.node1
  count       = local.node1_is_primary ? length(var.backend_servers) : 0

  name        = "${var.backend_servers[count.index].name}-svc"
  port        = var.backend_servers[count.index].port
  servername  = citrixadc_server.backend_node1[count.index].name
  servicetype = var.lb_service_type

  lbvserver   = citrixadc_lbvserver.app_lb_node1[0].name
}

resource "citrixadc_service" "backend_svc_node2" {
  provider    = citrixadc.node2
  count       = local.node2_is_primary ? length(var.backend_servers) : 0

  name        = "${var.backend_servers[count.index].name}-svc"
  port        = var.backend_servers[count.index].port
  servername  = citrixadc_server.backend_node2[count.index].name
  servicetype = var.lb_service_type

  lbvserver   = citrixadc_lbvserver.app_lb_node2[0].name
}
