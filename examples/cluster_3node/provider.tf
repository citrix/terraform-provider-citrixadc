terraform {
  required_providers {
    citrixadc = {
      source  = "citrix/citrixadc"
      version = "2.2.0"
    }
  }
}

# The provider endpoint must point to the Cluster IP (CLIP) address.
# It does not need to be reachable at init time (do_login defaults to false).
# The first real connection happens during the bootstrap poll, by which time
# the CLIP will have been created on the first node.
provider "citrixadc" {
  endpoint = "http://${var.clip}"
  username = var.adc_admin_username
  password = var.adc_admin_password
}
