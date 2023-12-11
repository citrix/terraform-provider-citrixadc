terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}

terraform {
  required_version = ">= 0.12"
}
provider "citrixadc" {
  endpoint = "http://172.16.149.100"
  username = var.username
  password = var.password
}
