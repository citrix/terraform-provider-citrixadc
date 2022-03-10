terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint = "http://<Citrix-ADC-Management-IP>"
  username = "<Citrix-ADC-Username>"
  password = "Citrix-ADC-Password"
}
