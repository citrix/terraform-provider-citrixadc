terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint = "http://10.222.74.157"
  username = "nsroot"
  password = "notnsroot"
}
