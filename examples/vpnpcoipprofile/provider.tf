terraform {
  required_providers {
    citrixadc = {
      source = "citrix/citrixadc"
    }
  }
}
provider "citrixadc" {
  endpoint = "http://loocalhost:8080"
  username = "<username>"
  password = "<password>"
}
