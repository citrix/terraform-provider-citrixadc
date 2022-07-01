terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint  = "http://localhost:8080"
  do_login  = true
  partition = "par1"
}
