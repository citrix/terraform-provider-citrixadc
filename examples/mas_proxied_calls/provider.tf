terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint   = "http://10.78.60.207"
  username   = "nsroot"
  password   = "nsroot"
  proxied_ns = "10.78.60.209"
}
