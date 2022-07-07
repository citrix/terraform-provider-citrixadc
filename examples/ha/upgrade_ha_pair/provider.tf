terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint = format("http://%s", var.primary_nsip)
  password = var.password
  alias = "primary"
}

provider "citrixadc" {
  endpoint = format("http://%s", var.secondary_nsip)
  password = var.password
  alias = "secondary"
}
