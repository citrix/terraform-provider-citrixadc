terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint = "http://52.10.11.134"
  username = "nsroot"
  password = "A3w5S9#13"
}