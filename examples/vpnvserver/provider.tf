terraform {
    required_providers {
        citrixadc = {
            source = "citrix/citrixadc"
        }
    }
}
provider "citrixadc" {
  endpoint = "http://localhost"
  username = "<username>"
  password = "<password>"
}
